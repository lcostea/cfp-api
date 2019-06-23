Concepts to present:
1. Pod and deployments
2. Readyness and liveness checks
3. Replicas, services, round robin, labels, selectors
4. Configuration: configmap and secret
5. Storage: persistent volumes and claims
6. Pod autoscaling: hpa
7. Multi container pods


kubectl run -i --tty busybox --image=busybox -- sh

Full qualified service name:
<service-name>.default.svc.cluster.local