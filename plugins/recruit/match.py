"""
Job Match Evaluator — Score candidate vs job description
Route: POST /api/recruit/match
"""

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        content = payload.get("content", "")
        job_desc = payload.get("job_description", "")

        llm = agent.llm if agent else None
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        scores = await llm.chat_json([
            {"role": "system", "content": "Evaluate candidate match 0-100. Return JSON: skills_match, experience_match, education_match, industry_match, overall_match. Add 80-word summary (English)."},
            {"role": "user", "content": f"Resume: {content}\nJob: {job_desc}"}
        ])

        envelop.payload = {
            **payload,
            "skills_match": scores.get("skills_match", 0),
            "experience_match": scores.get("experience_match", 0),
            "education_match": scores.get("education_match", 0),
            "industry_match": scores.get("industry_match", 0),
            "overall_match": scores.get("overall_match", 0),
            "summary": scores.get("summary", "")
        }
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop