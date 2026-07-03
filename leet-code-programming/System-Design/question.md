Interview Questions Asked in 30 LPA Job Interview

🚨 Question:
“An API fetches 1 million records from the database. Servers start failing. How do you prevent a total crash?”

This is not just a coding question.
This is a system design and scalability mindset test.

Here’s how I’d answer:
	1.	Pagination / Streaming Instead of Bulk Fetch
Never load 1M records into memory. Use pagination, cursor-based pagination, or streaming APIs to process data in chunks.

	2.	Database Optimization
Add proper indexes, avoid SELECT *, fetch only required columns, and use read replicas to reduce load on the primary database.

	3.	Backpressure & Rate Limiting
If downstream systems are slow, apply backpressure. Rate limit API requests to prevent sudden traffic spikes from crashing servers.

	4.	Caching Strategy
If data is frequently requested, cache it using Redis or in-memory caches to reduce database hits.

	5.	Asynchronous Processing
For large exports or analytics, use background jobs (Kafka, RabbitMQ, Spring Batch) instead of synchronous APIs.

	6.	Resource Limits & Autoscaling
Set JVM memory limits, use Kubernetes HPA for autoscaling, and add circuit breakers to prevent cascading failures.

	7.	Data Partitioning / Sharding
Split massive tables to avoid single-node bottlenecks.

Join my whatsapp channel : https://lnkd.in/dbQ3JzSw
