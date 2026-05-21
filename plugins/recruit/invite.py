"""
Interview Email Generator — Create invite or rejection email
Route: POST /api/recruit/invite
"""

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        decision = payload.get("decision", "hold")
        name = payload.get("name", "Candidate")

        llm = agent.llm if agent else None
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        email = await llm.chat_json([
            {"role": "system", "content": "Generate email: if strongly_recommend/recommend → interview invite (remote, 3 available time slots). Else → polite rejection. Return JSON: email_subject, email_body."},
            {"role": "user", "content": str(payload)}
        ])

        envelop.payload = {
            **payload,
            "email_subject": email.get("email_subject", ""),
            "email_body": email.get("email_body", "")
        }
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop