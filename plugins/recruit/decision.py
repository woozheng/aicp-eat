"""
Hiring Decision — Recommend hire, hold, or reject
Route: POST /api/recruit/decision
"""

async def execute(envelop, agent):
    try:
        payload = envelop.payload

        llm = agent.llm if agent else None
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        system_prompt = """
You are a professional HR recruitment analyst.
Based on candidate resume info, matching scores and interview questions, give final hiring decision.
Rules:
1. decision only choose one: strongly_recommend / recommend / hold / not_recommend
2. confidence: float number between 0 and 1
3. reasoning: exactly 3 detailed reasons, store as array
4. risk_factors: exactly 2 potential risk points, store as array
5. next_steps: exactly 3 follow-up action suggestions, store as array
All content must be written in English only.
Return standard JSON format only, no extra explanation.
"""

        decision_result = await llm.chat_json([
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": f"Full candidate evaluation data: {str(payload)}"}
        ])

        envelop.payload = {
            **payload,
            "decision": decision_result.get("decision", "hold"),
            "confidence": round(float(decision_result.get("confidence", 0.5)), 2),
            "reasoning": decision_result.get("reasoning", ["No valid reason provided"]),
            "risk_factors": decision_result.get("risk_factors", ["No obvious risk"]),
            "next_steps": decision_result.get("next_steps", ["Arrange further communication"])
        }
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop