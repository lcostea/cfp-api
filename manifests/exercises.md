### Prereq

Start the aks/minikube

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

### Multi container pods

TODO:--declarative spec, add one more container
https://matthewpalmer.net/kubernetes-app-developer/articles/multi-container-pod-design-patterns.html

### Readyness and liveness checks

`kubectl apply -f manifests/ready-live/go-api-demo-deploy-checks.yaml`

--increase timeout; add readyness

### Services
### Configuration: configmap and secret
### Pod autoscaling: hpa
### Ingress??


kubectl run -i --tty busybox --image=busybox -- sh

Full qualified service name:
<service-name>.default.svc.cluster.local