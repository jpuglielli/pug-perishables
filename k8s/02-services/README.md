# Services & networking (expose your app)

**Do (25 min):** Expose the nginx Deployment via a NodePort Service and reach it via curl/browser.

```bash
kubectl expose deployment nginx --type=NodePort --port=80
minikube service nginx --url   # then curl the URL
```

**Document (5 min):** Write up the port confusion you likely hit (port vs targetPort vs nodePort) and how selectors wired the Service to pods.

**Deliverable:** Reachable nginx endpoint + journal entry.
