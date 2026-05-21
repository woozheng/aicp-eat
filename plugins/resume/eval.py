import base64
import io
from pathlib import Path

def help():
    return {
        "route": "/api/resume/eval",
        "input": {"file": "base64 encoded file content", "filename": "resume.pdf"},
        "output": {"report": {...}, "scores": {...}, "suggestions": [...]},
        "description": "上传简历，自动生成完整评估报告"
    }

async def execute(envelop, agent):
    try:
        llm = agent.llm
        if not llm:
            envelop.payload = {"error": "LLM not available"}
            return envelop

        file_data = envelop.payload.get("file", "")
        filename = envelop.payload.get("filename", "resume.txt")
        
        # 解码文件内容
        try:
            content = base64.b64decode(file_data).decode("utf-8", errors="ignore")
        except:
            content = file_data
        
        # 提示词
        system_prompt = """你是一个专业的HR和职业发展顾问。请对简历进行全面评估。
        请从以下维度评估简历（满分100分）：
        1. 基本信息完整性 (20分)
        2. 教育背景 (15分)
        3. 工作经历 (25分)
        4. 技能展示 (20分)
        5. 项目经验 (20分)
        
        返回JSON格式：
        {
            "total_score": 总分,
            "grades": {
                "basic_info": {"score": 分数, "comment": "评价"},
                "education": {"score": 分数, "comment": "评价"},
                "experience": {"score": 分数, "comment": "评价"},
                "skills": {"score": 分数, "comment": "评价"},
                "projects": {"score": 分数, "comment": "评价"}
            },
            "strengths": ["优点1", "优点2", ...],
            "weaknesses": ["缺点1", "缺点2", ...],
            "suggestions": ["建议1", "建议2", ...],
            "summary": "总体评价"
        }"""
        
        result = await llm.chat_json([
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": f"请评估以下简历内容：\n\n{content[:8000]}"}
        ])
        
        envelop.payload = {
            "report": result,
            "filename": filename,
            "status": "success"
        }
    except Exception as e:
        envelop.payload = {"error": str(e), "status": "error"}
    return envelop