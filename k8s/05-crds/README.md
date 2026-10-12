# CRDs: the gateway to operators

**Do (25 min):** Define and apply your own CRD (e.g., `kind: Website`) and create a custom resource instance.

```bash
kubectl apply -f website-crd.yaml
kubectl apply -f website-instance.yaml
kubectl get websites
```

**Document (5 min):** Capture the key insight — a CRD only stores your resource; nothing acts on it yet. That gap is exactly what the operator will fill. This sets up Week 2.

**Deliverable:** Custom CRD + one instance registered + journal entry.
