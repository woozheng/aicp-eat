# core/__init__.py
"""AICP Protocol Engine — Three atomic units."""

import asyncio
import uuid
from typing import Dict, Optional

# ==============================================================================
# Agent — Mailbox that supports both dict and attribute access
# ==============================================================================
class Agent:
    """Agent mailbox — compatible with both agent.llm and agent.get('llm')."""
    
    def __init__(self, data: dict = None):
        if data:
            self.__dict__.update(data)
    
    def get(self, key, default=None):
        return self.__dict__.get(key, default)
    
    def __getitem__(self, key):
        return self.__dict__[key]
    
    def __setitem__(self, key, value):
        self.__dict__[key] = value
    
    def __contains__(self, key):
        return key in self.__dict__
    
    def __repr__(self):
        return f"Agent({self.__dict__})"


# ==============================================================================
# Registry — Global state
# ==============================================================================
agents: Dict[str, Agent] = {}
plugins: Dict[str, callable] = {}
bus: Optional["Bus"] = None
llm = None


# ==============================================================================
# Envelop — The letter of the information field
# ==============================================================================
class Envelop:
    __slots__ = (
        "sender", "receiver", "intent", "payload",
        "trace_id", "message_id", "channel_id", "ttl", "meta",
        "path_history", "original_initiator"
    )
    
    def __init__(self, sender="", receiver="", intent="", payload=None,
                 channel_id="", ttl=10, meta=None):
        self.sender = sender
        self.receiver = receiver
        self.intent = intent
        self.payload = payload if payload is not None else {}
        self.trace_id = f"tr_{uuid.uuid4().hex[:8]}"
        self.message_id = f"msg_{uuid.uuid4().hex[:6]}"
        self.channel_id = channel_id
        self.ttl = ttl
        self.meta = meta if meta is not None else {}
        self.path_history = []
        self.original_initiator = ""


# ==============================================================================
# Bus — Message scatterer
# ==============================================================================
class Bus:
    def __init__(self):
        self._subscriptions: Dict[str, set] = {}
    
    def subscribe(self, agent_id: str, channel: str):
        self._subscriptions.setdefault(channel, set()).add(agent_id)
    
    def unsubscribe(self, agent_id: str, channel: str):
        if channel in self._subscriptions:
            self._subscriptions[channel].discard(agent_id)
    
    async def publish(self, channel: str, envelop: Envelop):
        if not channel:
            return
        
        if channel.startswith("grp."):
            for sub in self._subscriptions.get(channel, set()):
                if sub == envelop.sender:
                    continue
                env = Envelop(
                    sender=envelop.sender, receiver=sub,
                    intent=envelop.intent, payload=dict(envelop.payload),
                    channel_id=channel, ttl=envelop.ttl, meta=dict(envelop.meta)
                )
                env.path_history = list(envelop.path_history)
                env.original_initiator = envelop.original_initiator or envelop.sender
                asyncio.create_task(engine_route(env))
        else:
            envelop.receiver = channel
            asyncio.create_task(engine_route(envelop))


# ==============================================================================
# Router — Blind router
# ==============================================================================
async def engine_route(envelop: Envelop):
    if envelop.ttl <= 0:
        return
    envelop.ttl -= 1
    
    # 0. Plugin direct call
    if envelop.receiver in plugins:
        plugin = plugins[envelop.receiver]
        http_agent = Agent({
            "id": "http",
            "workflow": [],
            "llm": llm,
            "bus": bus,
        })
        result = await plugin(envelop, http_agent)
        return result
    
    # 1. Agent route
    agent = agents.get(envelop.receiver)
    if not agent:
        return  # DEAD
    
    if envelop.receiver == envelop.sender:
        return
    
    saved_receiver = envelop.receiver
    saved_channel = envelop.channel_id
    envelop.receiver = ""
    
    out = await workflow_run(envelop, agent)
    if not out:
        return  # DEAD
    
    if not out.receiver:
        if out.path_history:
            out.receiver = out.path_history.pop()
        elif saved_channel and saved_channel.startswith("grp."):
            out.receiver = saved_channel
        else:
            out.receiver = out.original_initiator or saved_receiver
    
    out.sender = agent.get("id", "")
    await bus.publish(out.receiver, out)


# ==============================================================================
# Workflow — Plugin chain executor
# ==============================================================================
async def workflow_run(envelop: Envelop, agent: Agent) -> Optional[Envelop]:
    steps = agent.get("workflow", [])
    current = envelop
    
    for name in steps:
        plugin = plugins.get(name)
        if not plugin:
            continue
        current = await plugin(current, agent)
        if not current:
            return None  # DEAD
    
    return current