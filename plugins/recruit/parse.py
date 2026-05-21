"""
Resume Parser — Extract structured info from resume text
Route: POST /api/recruit/parse
"""

async def execute(envelop, agent):
    try:
        payload = envelop.payload
        content = payload.get("content", "")

        llm = agent.llm if agent else None
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        parsed = await llm.chat_json([
            {"role": "system", "content": "Extract structured resume info. Return JSON only: name, email, phone, years_of_experience, skills (array), education, current_company. Use English field values."},
            {"role": "user", "content": content}
        ])

        envelop.payload = {
            "name": parsed.get("name", ""),
            "email": parsed.get("email", ""),
            "phone": parsed.get("phone", ""),
            "years_of_experience": parsed.get("years_of_experience", 0),
            "skills": parsed.get("skills", []),
            "education": parsed.get("education", ""),
            "current_company": parsed.get("current_company", ""),
            "content": content
        }
        return envelop

    except Exception as e:
        envelop.payload = {"error": str(e)}
        return envelop