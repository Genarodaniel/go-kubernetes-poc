#to build the docker image
docker build -t danielsgenaro/address-crud-1:latest -f Dockerfile .

#to push the image to docker registry
docker push danielsgenaro/address-crud-1:latest

#to run an image
docker run --rm -p 8080:8080 danielsgenaro/address-crud-1:latest

#to delete the cluster
kind delete cluster --name=address-crud-1

#to create the cluster
kind create cluster --name=address-crud-1

#to get clusters
kind get clusters

#to verify if the cluster is created and get the info
kubectl cluster-info --context kind-address-crud-1

#get nodes
kubectl get nodes

#get nodes
kubectl get svc

#apply the yaml to our running pods
kubectl apply -f kubernetes/deployment.yaml
kubectl apply -f kubernetes/service.yaml


#to forward the port from local to the svc
kubectl port-forward svc/serversvc 8080:8080

#describe pod
kubectl describe pod server-7fd9c997d7-9qjs2

#SVC type: default = port forward
#SVC type: type: LoadBalancer = create an external IP for your pod