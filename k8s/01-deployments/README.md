# Core objects & your first deployment

**Do (25 min):** On a local cluster (minikube, kind, or Docker Desktop), deploy an nginx Deployment with 3 replicas and confirm it's running.

```bash
kubectl create deployment nginx --image=nginx --replicas=3
kubectl get pods -o wide
```

**Document (5 min):** Log what went well, what broke, and how you fixed it. Note the mental model of Deployment → ReplicaSet → Pods.

**Deliverable:** 3 nginx pods running, captured in your journal.
