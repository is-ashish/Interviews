# 🚨 CrashLoopBackOff in Kubernetes

A pod stuck in **CrashLoopBackOff** means Kubernetes is repeatedly trying to start the pod, but it keeps failing.

---

## 🔎 Debugging Steps

1. **Check pod status and events**
   ```bash
   kubectl describe pod <pod-name>

Look for events like OOMKilled, ImagePullBackOff, or failing probes.


Inspect logs
kubectl logs <pod-name> --previous

--previous shows logs from the last terminated container.

Check container image
Verify the image exists, is accessible, and not corrupted.

Ensure registry credentials are correct if private.
Validate configuration

Look at environment variables, ConfigMaps, Secrets.

Check if required files or volumes are mounted correctly.
Probe settings

Misconfigured liveness/readiness probes can kill the pod prematurely.

Resource limits
If CPU/memory requests/limits are too strict, the pod may get evicted or OOMKilled.
Node health
Check if the node has enough resources or if there are taints preventing scheduling.
---
## ⚠️ Possible Causes of CrashLoopBackOff

### Application-level issues
- Bug in code causing immediate crash.
- Missing dependencies or misconfigured startup commands.

### Configuration errors
- Wrong environment variables, secrets, or config files.
- Incorrect command/entrypoint in Dockerfile.

### Resource constraints
- Pod exceeds memory → OOMKilled.
- CPU throttling causing instability.

### Image problems
- Bad build, missing binaries, or wrong tag.
- Authentication failure with private registry.

### Probe misconfiguration
- Liveness probe too aggressive → container killed before it stabilizes.

### Cluster/node issues
- Disk pressure, insufficient memory, or taints preventing pod from running.

---
🧠 Interview Tip
When asked this, structure your answer:
Step 1: “I’d start by checking logs and events with kubectl describe and kubectl logs.”
Step 2: “Then I’d validate configuration, probes, and resource limits.”
Step 3: “Finally, I’d consider application bugs, image issues, or cluster resource problems.”
This shows you can systematically debug rather than guessing.

---

Would you like me to also add a **ready‑to‑use “cheat sheet” version** (super concise bullet points only) so you can quickly glance at it before interviews?
