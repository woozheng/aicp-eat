"""
Interview Question Generator — Generate 5 structured questions
Route: POST /api/recruit/questions
"""

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        content = payload.get("content", "")
        job_desc = payload.get("job_description", "")
        skills = payload.get("skills", [])

        llm = agent.llm if agent else None
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        questions = await llm.chat_json([
            {"role": "system", "content": "Generate 5 interview questions: 3 technical (skill gaps), 1 behavioral, 1 situational. Return JSON array: [{type: technical/behavioral/situational, question}]."},
            {"role": "user", "content": f"Resume: {content}\nJob: {job_desc}\nSkills: {skills}"}
        ])

        envelop.payload = {
            **payload,
            "questions": questions
        }
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop