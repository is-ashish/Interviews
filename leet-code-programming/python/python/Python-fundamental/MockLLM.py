from fastapi import FastAPI, HTTPException
from pydantic import BaseModel, Field
import time
import uuid
import asyncio

app = FastAPI()


class UserRequest(BaseModel):
    user_id:str=Field(...,min_length=1)
    session_id:str=Field(...,min_length=1)
    query:str=Field(..., min_length=1)


async def MockLLM(question:str):
    await asyncio.sleep(3)
    for ch in question:
        int(ch)
    
    return {
        "question": question,
        "answer": f"This is mocked answer",
    }
    
@app.post("/chatting")
async def getAnswer(requestBody:UserRequest):
    try:
        reqID = str(uuid.uuid4())
        startTime = time.perf_counter()
        res = await MockLLM(requestBody.query)
        endTime =  time.perf_counter()
        latency_ms = (endTime-startTime)*1000
        return {
            "answer": res["answer"],
            "Latency": f"{latency_ms:.2f}",
            "Request_ID": reqID
        }
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Internal Server Error {str(e)}")