def help():
    return {
        "route": "/api/summarize",
        "input": {"content": "long text to summarize"},
        "output": {"summary": "3 sentence summary", "keywords": ["keyword1", "keyword2", ...]},
        "description": "Summarize text into 3 sentences with keywords"
    }

async def execute(envelop, agent):
    try:
        llm = agent.llm
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        text = envelop.payload.get("content", "")
        if not text or len(text.strip()) < 50:
            envelop.payload = {"error": "Text too short, minimum 50 characters"}
            return envelop

        lang = envelop.payload.get("lang", "en")

        system_prompt = """You are a professional text summarizer. Provide a concise summary following these rules:
1. Extract exactly 3 key points and write each as a complete sentence
2. Extract 5-8 relevant keywords
3. Respond in JSON format: {"summary": ["sentence1", "sentence2", "sentence3"], "keywords": [...]}
4. Keywords should be single words or short phrases (2-3 words max)
5. Focus on the main topics, entities, and actions in the text"""

        if lang == "zh":
            system_prompt = """你是一个专业的文本摘要生成器。请按以下规则提供摘要：
1. 提取3个关键点，每个写成完整句子
2. 提取5-8个相关关键词
3. 用JSON格式回复：{"summary": ["句子1", "句子2", "句子3"], "keywords": [...]}
4. 关键词应是单词或短词组（最多2-3个词）
5. 关注文本的主要主题、实体和行为"""

        result = await llm.chat_json([
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": text}
        ])

        summary_text = ". ".join(result.get("summary", []))
        if lang == "zh":
            summary_text = "。".join(result.get("summary", []))

        envelop.payload = {
            "summary": summary_text,
            "keywords": result.get("keywords", []),
            "lang": lang
        }
    except Exception as e:
        envelop.payload = {"error": str(e)}
    return envelop