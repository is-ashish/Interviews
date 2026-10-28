from fastapi import FastAPI, HTTPException
from pydantic import BaseModel, Field
from contextlib import contextmanager
import sqlite3
import asyncio
import time
import uuid

app = FastAPI()


@contextmanager
def DB_Connection(db_name):
    db = sqlite3.connect(db_name)
    try:
        yield db # provide connection
        db.commit() # execute table
    except Exception as e:
        db.rollback()
        print("DB Error --: ", str(e))
    finally:
        db.close()
        print("Connection Close")

class UserRequest(BaseModel):
    question:str=Field(..., description="User query", min_length=1)
    user_id: str=Field(...,description="User Id")
    session_id: str=Field(...,description="Session Id")
    
async def mockLLM(query:str):
    await asyncio.sleep(3)
    return {
        "question":query,
        "answer":f"This is Mocked LLM response about - {query}"
    }
    
    
@app.get("/welcome")
def welcome():
    return {
        "message": "Welcome To Mock LLM Application"
    }

@app.post("/query")
async def PostQuestion(reqBody:UserRequest) -> dict:
    
    try:
        requestID = str(uuid.uuid4())
        startTime = time.time()
        result =await mockLLM(reqBody.question)
        endTimeTime = time.time()
        latency= f"{endTimeTime-startTime:.8f}"
        
        with DB_Connection("test_db") as connection:
            cursor = connection.cursor()
            cursor.execute("CREATE TABLE IF NOT EXISTS chatting(Request_Id TEXT PRIMARY KEY, User_ID TEXT, Questions TEXT, Answer TEXT, Latency TEXT)")
            cursor.execute("INSERT INTO chatting VALUES (?, ?, ?, ?, ?)", (requestID,reqBody.user_id, reqBody.question, result["answer"],latency))
            row = cursor.execute("select * from chatting where Request_Id=?", (requestID,))
            row = row.fetchone()
            if row:
                response = {
                    "Request_Id": row[0],
                    "User_ID": row[1],
                    "Questions": row[2],
                    "Answer": row[3],
                    "Latency":row[4],
                }
            return response
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"{str(e)}")
       
       
@app.get("/all_chat/{user_id}")
def all_chatting(user_id:str) -> list[dict]:
    response = []
    with DB_Connection("test_db") as connection:
        cursor = connection.cursor()
        cursor.execute("select * from chatting where User_ID=?", (user_id,))
        rows =  cursor.fetchall()
        
        if not rows:
            raise HTTPException(status_code=404, detail=f"No chat found for user ID - {user_id}")
        
        for row in rows:
            response.append({
                "Request_Id": row[0],
                    "User_ID": row[1],
                    "Questions": row[2],
                    "Answer": row[3],
                    "Latency":row[4],
                })
    return response
        
                
            
            
    
        
    
    
    
