import aiohttp
import asyncio

async def check_service(session, address):
    url = f"http://{address}/status"
    try:
        async with session.get(url, timeout=0.5) as response:
            return response.status == 200
    except Exception as e:
        return e

async def health_check(addresses):
    results = {}
    async with aiohttp.ClientSession() as session:
        tasks = [check_service(session, addr) for addr in addresses]
        statuses = await asyncio.gather(*tasks)
        for addr, status in zip(addresses, statuses):
            results[addr] = status
    return results

# Example usage
async def main():
    services = ["localhost:8000", "localhost:8080", "enterprise:8008"]
    result = await health_check(services)
    print(result)

if __name__ == "__main__":
    asyncio.run(main())
