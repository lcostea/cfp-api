### Prereq

AKS API Server:
https://cniasi-workshop-june-dns-ae480410.hcp.northeurope.azmk8s.io


Current context

`kubectl config use context minikube`

OR

`kubectx minikube`

Namespaces

`kubectl create ns liviucostea`


`kubens liviucostea`

OR 

`kubectl config set-context $(kubectl config current-context) --namespace=liviucostea`

### Pod and deployments

`kubectl apply -f manifests/deployment/go-api-demo-pod.yaml`

`kubectl run -it busybox --image=busybox --restart=Never /bin/sh`

`wget -S http://<pod-ip>:3000/callforpapers -O /dev/null`


Upgrade to a deployment, includes a replicaset


`kubectl apply -f manifests/deployment/go-api-demo-deploy.yaml`

TODO: make a rollout for version 0.1.1 then rollback to version 0.1.0


### Readyness and liveness checks

`kubectl apply -f manifests/ready-live/go-api-demo-deploy-checks.yaml`

--increase timeout; add readyness

### Services

Use the checks deployment from above

`kubectl apply -f ./manifests/ready-live/go-api-demo-deploy-checks.yaml`

`k apply -f manifests/service/go-api-demo-service.yaml`

`kubectl run -it busybox --image=busybox --restart=Never /bin/sh`

--cleanup busybox


Full qualified service name:
<service-name>.default.svc.cluster.local

TODO: create the service, apply the busybox and then make some queries and see they are going to different pods

### Configuration: configmap and secret



### Multi container pods

Git-sync sidecar

`kubectl apply -f manifests/deployment/go-api-demo-multi.yaml`

`kubectl exec -it <multi-deploy-pod-name> /bin/sh`



### Pod autoscaling: hpa

Metrics server and heapster are already installed on AKS
`kubectl get deployments --all-namespaces`

### Ingress??



