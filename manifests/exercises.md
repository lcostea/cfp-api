## Prereq

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

## Pod and deployments

`kubectl apply -f manifests/deployment/go-api-demo-pod.yaml`

`kubectl run -it busybox --image=busybox --restart=Never /bin/sh`

`wget -S http://<pod-ip>:3000/callforpapers -O /dev/null`


Upgrade to a deployment, includes a replicaset


`kubectl apply -f manifests/deployment/go-api-demo-deploy.yaml`


## Readyness and liveness checks

`kubectl apply -f manifests/ready-live/go-api-demo-deploy-checks.yaml`

--increase timeout; add readyness

## Services

Use the checks deployment from above

`kubectl apply -f ./manifests/ready-live/go-api-demo-deploy-checks.yaml`

`k apply -f manifests/service/go-api-demo-service.yaml`

`kubectl run -it busybox --image=busybox --restart=Never /bin/sh`

--cleanup busybox

`kubectl port-forward svc/go-api-demo 8080:80`

try to access:
http://localhost:8080/callforpapers


Fully qualified service name:
go-api-demo.default.svc.cluster.local

## Configuration: configmap and secret

create the deployment, pod will give config error:

`kubectl apply -f manifests/config/go-api-demo-deploy-config.yaml`

creates the cm:

`kubectl apply -f manifests/config/go-api-demo-cm.yaml`

## Rollout and rollback

modify image in deploy/go-api-demo-config for 0.1.1 and apply
`kubectl get replicaset`

`kubectl rollout status deploy/go-api-demo-config`

`kubectl rollout history deploy/go-api-demo-config`

undo to the previous deploy:

`kubectl rollout undo deploy/go-api-demo-config`
`kubectl get replicaset`

## Multi container pods

Git-sync sidecar

`kubectl apply -f manifests/deployment/go-api-demo-multi.yaml`

`kubectl exec -it <multi-deploy-pod-name> /bin/sh`


## Pod autoscaling: hpa

Metrics server and heapster are already installed on AKS
`kubectl get deployments --all-namespaces`


## Ingress

After creating an ingress controller (like nginx), then we can create ingress rules
https://kubernetes.io/docs/concepts/services-networking/ingress/



