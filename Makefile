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

#describe ingress

kubectl describe ingress address-ingress


#SVC type: default = port forward
#SVC type: type: LoadBalancer = create an external IP for your pod


#kubeconfig aws 
aws eks --region us-east-1 update-kubeconfig --name address-crud-1

aws eks list-clusters

eksctl utils associate-iam-oidc-provider --region us-east-1 --cluster address-crud-1 --approve

#verify the provider
aws eks describe-cluster --name address-crud-1 --query "cluster.identity.oidc.issuer" --output text

#download IAM Policy of load balancer
curl -o iam-policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/main/docs/install/iam_policy.json

#creating the IAM Policy
aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy --policy-document file://iam-policy.json
#OUTPUT: 
# {
#     "Policy": {
#         "PolicyName": "AWSLoadBalancerControllerIAMPolicy",
#         "PolicyId": "ANPAXZ5NF77557TY5U4CG",
#         "Arn": "arn:aws:iam::536697241595:policy/AWSLoadBalancerControllerIAMPolicy",
#         "Path": "/",
#         "DefaultVersionId": "v1",
#         "AttachmentCount": 0,
#         "PermissionsBoundaryUsageCount": 0,
#         "IsAttachable": true,
#         "CreateDate": "2025-01-26T18:26:51+00:00",
#         "UpdateDate": "2025-01-26T18:26:51+00:00"
#     }
# }

#Creating kubernetes service account and associating with the IAM
eksctl create iamserviceaccount \
  --cluster address-crud-1 \
  --namespace kube-system \
  --name aws-load-balancer-controller \
  --attach-policy-arn arn:aws:iam::536697241595:policy/AWSLoadBalancerControllerIAMPolicy \
  --approve

# adding the eks chart to helm app 
helm repo add eks https://aws.github.io/eks-charts
helm repo update

#Installing the helm chart
helm install aws-load-balancer-controller eks/aws-load-balancer-controller \
  -n kube-system \
  --set clusterName=address-crud-1 \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-load-balancer-controller \
  --set region=us-east-1 \
  --set vpcId=vpc-0f04b6fc83e2c6b29



kubectl expose deployment server  --type=LoadBalancer  --name=serversvc

kubectl get service/serversvc |  awk {'print $1" " $2 " " $4 " " $5'} | column -t

kubectl describe ingress address-ingress -n default

kubectl port-forward svc/serversvc 8080:8080 -n default