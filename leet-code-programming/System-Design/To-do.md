(1) use read replicas to reduce load on the primary database.
(2) Resource Limits & Autoscaling -: use Kubernetes HPA for autoscaling, and add circuit breakers to prevent cascading failures.
(3) Sharding and connection pooling
(4) • Circuit Breakers: Prevent cascading failures if one service is down.

