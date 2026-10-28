Great question — asyncio often comes up in senior Python interviews because it tests your understanding of concurrency, event loops, and when to use async vs threads vs multiprocessing.
---
🔑 Frequently Asked Interview Questions on `asyncio`
Core Concepts
What is asyncio and how does it differ from threading or multiprocessing?
Explain the event loop in Python. How does it schedule tasks?
What’s the difference between await and async keywords?
How does asyncio.gather() differ from asyncio.create_task()?
What happens if you call a blocking function inside an async coroutine?
Practical Usage
How would you design a high‑throughput API using FastAPI or aiohttp with asyncio?
Show how you’d run multiple I/O tasks concurrently (e.g., fetching data from multiple APIs).
How do you handle exceptions in async tasks?
How do you integrate asyncio with existing synchronous libraries?
Advanced Topics
How does the Global Interpreter Lock (GIL) affect asyncio?
When would you prefer asyncio over threads or multiprocessing?
How do you cancel tasks gracefully in asyncio?
What’s the difference between cooperative multitasking (asyncio) and pre‑emptive multitasking (threads)?
---
📊 Example Interview Exercise
Question:
Write a Python function that fetches data from three APIs concurrently and returns the results.
Answer (simplified):
import asyncio
import aiohttp

async def fetch(session, url):
    async with session.get(url) as response:
        return await response.text()

async def main():
    urls = ["https://api1.com", "https://api2.com", "https://api3.com"]
    async with aiohttp.ClientSession() as session:
        results = await asyncio.gather(*(fetch(session, u) for u in urls))
    print(results)

asyncio.run(main())

---
👉 In interviews, the key is not just writing code but explaining when and why you’d use asyncio — e.g., for I/O‑bound tasks like network calls, not CPU‑bound tasks (where multiprocessing is better).
Would you like me to prepare a mock interview Q&A set specifically on Python concurrency (asyncio vs threads vs multiprocessing) so you can practice answering them fluently? That’s a common senior‑level test.