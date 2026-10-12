# Journal

### Day 1

Getting the service and deployments configured was pretty straight forward. Useful to know that you can use kubectl explain to get documentation on a resource. 

Could not start minikube, because of a zscalar issue, will update here when resolved.

**Update — resolved.** There were actually two problems:

- **Broken cluster:** the old cluster's `/etc/kubernetes` was empty on restart, so the kubelet couldn't bootstrap and the apiserver never came up (`K8S_APISERVER_MISSING`). Fix: `minikube delete` and start fresh.
- **Zscaler:** it intercepts TLS to `registry.k8s.io`, so pulls failed with `unable to get local issuer certificate`. Fix: export the Zscaler Root CA from the System keychain into `~/.minikube/certs/`. minikube installs it into the node on start.

  ```sh
  security find-certificate -a -c "Zscaler Root CA" -p /Library/Keychains/System.keychain > ~/.minikube/certs/zscaler-root-ca.pem
  ```

- Saved defaults with `minikube config set`: docker driver, containerd, 4 CPUs, 8 GB.
- The "Failing to connect to registry.k8s.io" warning on start is harmless; pulls work.
- Old images like `echoserver:1.4` won't pull on containerd 2.x (v1 manifest). Use `registry.k8s.io/e2e-test-images/agnhost` instead.
- Docker Desktop was using 474 GB of disk (~395 GB reclaimable). Prune it at some point.