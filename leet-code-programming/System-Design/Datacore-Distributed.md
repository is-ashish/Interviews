1. Data inconsistency in distributed storage
Q: A customer reports stale reads in a replicated volume. What do you do?
A: Verify consistency model → check quorum & replication → inspect leader election & logs → detect network partitions → fix misconfiguration (quorum/async writes) → enforce strong consistency if needed → add observability & tests.
 
2. High latency with NVMe
Q: NVMe-backed system shows high latency under load. Why?
A: Profile I/O path → identify syscall/lock bottlenecks → use eBPF/perf → optimize with SPDK, poll-mode drivers, zero-copy → increase queue depth → validate with fio benchmarks.
 
3. CSI mount failures after node restart
Q: StatefulSet volumes fail to mount intermittently.
A: Check CSI logs & kubelet → verify NodeStage/Publish → detect stale attachments → ensure idempotency → fix node identity issues → add retries & resilience.
 
4. Ambiguous system design
Q: Design a high-performance cloud-native storage engine.
A: Clarify requirements → separate control/data plane → use Raft → define replication model → integrate with Kubernetes via CSI → evaluate CAP trade-offs.
 
5. Production crash under load
Q: Data plane crashes in production.
A: Mitigate impact → analyze logs/core dumps → check race conditions/memory leaks → fix concurrency issues → deploy hotfix → write RCA + tests.
 
6. Architecture disagreement
Q: Kernel vs user-space storage debate.
A: Compare performance vs maintainability → run benchmarks → align with product goals → document decision → ensure team consensus.
 
7. Unhappy enterprise customer
Q: Customer says system is unreliable.
A: Gather data → identify misconfigurations → explain clearly → fix architecture → feed learnings into product improvements.
 
8. Open-source performance issue
Q: Found major issue affecting community.
A: Reproduce → benchmark → create issue + PR → engage community → ensure backward compatibility.
 
9. Scaling failure (50 → 500 nodes)
Q: System breaks at scale.
A: Identify bottlenecks (etcd/control plane) → shard metadata → reduce centralization → run large-scale tests.
 
10. Mentoring engineer
Q: Junior struggling with distributed systems.
A: Simplify concepts → pair program → assign focused tasks → build fundamentals → encourage design thinking.
 
🚀 OpenEBS-Specific Scenario Q&A
11. OpenEBS volume goes read-only
 
Q: A volume suddenly becomes read-only in OpenEBS.
A:
 
Check replica health (Jiva/cStor/Mayastor)
 
Verify quorum loss
 
Inspect pool/node failures
 
Check disk pressure or I/O errors
 
Restore replicas or rebalance
 
Ensure proper replication factor
 
12. Mayastor performance issue
 
Q: High latency in Mayastor volumes.
A:
 
Inspect SPDK usage
 
Check CPU pinning & hugepages
 
Validate NVMe performance
 
Tune poll-mode threads
 
Analyze I/O queue depth
 
13. Replica rebuild is slow
 
Q: Rebuild of OpenEBS replica is taking too long.
A:
 
Check network bandwidth
 
Validate snapshot chain length
 
Optimize rebuild strategy (incremental vs full)
 
Monitor disk throughput
 
Tune parallel rebuilds
 
14. Volume stuck in “Pending”
 
Q: PVC not getting bound in OpenEBS.
A:
 
Check StorageClass config
 
Verify pool availability
 
Inspect operator logs
 
Check node labels & taints
 
Ensure sufficient capacity
 
15. Data loss after node failure
 
Q: Node failure caused data loss. Why?
A:
 
Insufficient replication factor
 
Replicas scheduled on same failure domain
 
No anti-affinity rules
 
Fix by enforcing topology-aware scheduling
 
16. CSI driver crash loop
 
Q: OpenEBS CSI pods are crashing.
A:
 
Check logs for API errors
 
Verify CRDs
 
Check Kubernetes version compatibility
 
Validate RBAC permissions
 
17. Split-brain scenario
 
Q: Two replicas diverge in OpenEBS.
A:
 
Check quorum enforcement
 
Inspect network partitions
 
Validate leader election
 
Add fencing mechanisms
 
18. High CPU usage in storage pods
 
Q: Storage engine consuming high CPU.
A:
 
Profile with eBPF
 
Check busy polling (SPDK)
 
Optimize thread allocation
 
Reduce unnecessary loops
 
19. Frequent volume detaches
 
Q: Volumes keep detaching/attaching.
A:
 
Check node stability
 
Inspect kubelet restarts
 
Validate CSI timeouts
 
Fix node health issues
 
20. Customer wants multi-zone HA
 
Q: Design OpenEBS for multi-zone resilience.
A:
 
Use topology-aware scheduling
 
Spread replicas across zones
 
Configure anti-affinity
 
Ensure quorum across zones
 
Test zone failure scenarios