import aiohttp
import asyncio

async def health_check(addresses):
    result={}
    
    async def check_one(session, address):
        url = address
        try:
            async with session.get(url,timeout=aiohttp.ClientTimeout(total=0.5)) as response:
                result[address] = (response.status == 200)
        except Exception:
            # Timeout, connection error, etc. → unhealthy
            result[address] = False
            print("Exception Occured")
                
        
    async with aiohttp.ClientSession() as session:
        task = [check_one(session, address) for address in addresses]
        print(task)
        await asyncio.gather(*task)
    
    return result
        

results = asyncio.run(health_check(["https://google.com", "https://python.org", "http://badurl.test"]))

print(results)


