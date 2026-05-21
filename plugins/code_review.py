async def execute(envelop, agent):
    try:
        # 获取LLM实例
        llm = agent.llm
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        # 获取用户提交的代码内容
        code_content = envelop.payload.get("content", "")
        if not code_content.strip():
            envelop.payload = {"error": "请提交需要审查的代码片段"}
            return envelop

        # 代码审查核心逻辑，使用JSON格式返回结构化结果
        review_result = await llm.chat_json([
            {
                "role": "system",
                "content": """你是专业的代码审查助手，对用户提交的代码进行全面审查，严格按照JSON格式返回结果。
审查维度：
1. 代码质量：命名规范、可读性、代码规范、注释完整性
2. 安全漏洞：SQL注入、XSS、敏感信息泄露、权限问题、不安全的函数使用
3. 性能问题：冗余代码、无效循环、内存占用、低效算法
4. 改进建议：具体可落地的优化方案
5. 综合评分：0-100分

返回格式：
{
    "score": 分数,
    "quality_analysis": "代码质量分析",
    "security_risks": ["漏洞1", "漏洞2"],
    "performance_issues": ["问题1", "问题2"],
    "improvement_suggestions": ["建议1", "建议2"],
    "optimized_code": "优化后的代码（无优化则为空）"
}"""
            },
            {"role": "user", "content": f"请审查以下代码：\n{code_content}"}
        ])

        # 赋值结果并返回
        envelop.payload = review_result
    except Exception as e:
        envelop.payload = {"error": f"代码审查失败：{str(e)}"}
    return envelop