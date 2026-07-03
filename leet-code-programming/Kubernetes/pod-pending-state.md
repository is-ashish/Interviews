
When a Kubernetes Pod is stuck in the **Pending** state, it means the scheduler has accepted the Pod but cannot place it on any node. The containers haven’t started yet. This is usually a scheduling problem rather than a runtime one.

---

## 🔎 How to Troubleshoot

1. **Check Pod Events**
   ```bash
   kubectl describe pod <pod-name>

Look at the Events section for clues (e.g., “0/3 nodes are available: insufficient memory”).
Check Node Resources
kubectl get nodes
kubectl describe node <node-name>

Verify CPU, memory, and disk availability. Pods may be pending if nodes don’t have enough resources.
Check Node Selectors / Affinity
If the Pod has nodeSelector, affinity, or taints/tolerations, ensure there are nodes that match.
Misconfigured selectors often cause Pending state.
Check Storage
If the Pod requests a PersistentVolumeClaim (PVC), ensure a matching PersistentVolume (PV) exists and is bound.
Pending PVCs will keep Pods pending.
Check Quotas and Limits
kubectl describe quota -n <namespace>

Namespace resource quotas or Pod limits may prevent scheduling.
Check Network Policies
Rarely, restrictive policies can block scheduling if init containers or sidecars depend on connectivity.
---
⚠️ Common Causes
1. Cause	Symptom in Events / Logs
2. Insufficient CPU/Memory	“0/3 nodes are available: insufficient CPU”
3. NodeSelector mismatch	“no nodes match pod node affinity/selector”
4. Taints/Tolerations mismatch	“pod didn’t tolerate node taint”
5. PVC not bound	“pod has unbound immediate PersistentVolumeClaim”
6. ßResource quota exceeded	“exceeded quota: cpu, memory”
---
✅ Next Steps
Start with kubectl describe pod to see why scheduling failed.
If it’s a resource issue, scale your cluster or adjust requests/limits.
If it’s a selector/taint issue, fix the Pod spec or node labels.
If it’s storage, ensure PVCs are properly bound.
---
In short:
Pending = scheduler can’t find a suitable node. The fix depends on whether the bottleneck is resources, configuration, or storage.

---

Would you like me to extend this into a **step-by-step checklist with commands and expected outputs** so you can quickly pinpoint the exact reason for a Pod being Pending?
