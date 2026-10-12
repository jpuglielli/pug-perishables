# State, health & packaging

**Do (25 min):** Add liveness + readiness probes to your Deployment, then install one app via Helm.

```bash
helm install my-nginx bitnami/nginx
kubectl describe pod <pod>   # verify probes
```

**Document (5 min):** Record probe behavior you observed (restarts, readiness gating) and your first impressions of Helm vs. raw YAML.

**Deliverable:** Probes working + one Helm release installed + journal entry.
