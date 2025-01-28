# Project Setup Guide

This guide will walk you through the process of setting up and deploying your application using Kubernetes, Terraform, and Docker. Please ensure you have the following prerequisites installed on your system:

- [Terraform](https://www.terraform.io/downloads.html)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [AWS CLI](https://aws.amazon.com/cli/)
- [Docker](https://docs.docker.com/get-docker/) (for building and running containers)
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) (for local Kubernetes setup)
- [Docker Compose](https://docs.docker.com/compose/) (for running locally)

---

## Running Locally

### 1. Using Docker Compose
To build and run the Docker images locally, use Docker Compose:
```bash
docker-compose up --build
```
This will build the Docker images as defined in your `docker-compose.yml` file and start the containers.

### 2. Using kind
To test your Kubernetes deployment locally using kind (Kubernetes in Docker):

#### Create a kind Cluster
```bash
kind create cluster
```

#### Deploy to kind
Apply the Kubernetes deployment configuration:
```bash
kubectl apply -f kubernetes/deployment.yaml
```

#### Update the Image in kind
If you need to update the Docker image being used in your deployment:
1. Build the updated Docker image:
   ```bash
   docker build -t <your-dockerhub-username>/<your-app-name>:<new-tag> .
   ```
   Replace `<your-dockerhub-username>`, `<your-app-name>`, and `<new-tag>` with the appropriate values.

2. Update the `image` field in `kubernetes/deployment.yaml` with the new image tag:
   ```yaml
   image: <your-dockerhub-username>/<your-app-name>:<new-tag>
   ```

3. Apply the updated deployment:
   ```bash
   kubectl apply -f kubernetes/deployment.yaml
   ```

#### Access the Application
Since LoadBalancers are not supported natively in kind, use the following command to access your service:
```bash
kubectl port-forward service/<service-name> <local-port>:<service-port>
```
Replace `<service-name>`, `<local-port>`, and `<service-port>` with the appropriate values for your service.

---

## Running in the Cloud

### 1. Set Up Infrastructure with Terraform
Navigate to the `terraform` directory and initialize Terraform:
```bash
cd terraform
terraform init
```

Apply the Terraform configuration to create the Kubernetes service:
```bash
terraform apply
```
You will be prompted to enter your AWS Access Key and Secret Key during this process. Provide them as requested.

> **Note:** Review the Terraform execution plan and type `yes` to confirm the changes.

#### Destroy Infrastructure with Terraform
If you need to tear down the infrastructure created by Terraform, use the following command:
```bash
terraform destroy
```
This will remove all resources created by the Terraform configuration. Review the plan and type `yes` to confirm.

### 2. Build Docker Images for Server Deployment
To build the Docker image directly for server deployment, run the following command:
```bash
docker build -t <your-dockerhub-username>/<your-app-name>:<tag> .
```
Replace `<your-dockerhub-username>`, `<your-app-name>`, and `<tag>` with your Docker Hub username, application name, and desired tag.

### 3. Push the Docker Image to Docker Hub
Login to Docker Hub:
```bash
docker login
```
Push the image to Docker Hub:
```bash
docker push <your-dockerhub-username>/<your-app-name>:<tag>
```

### 4. Deploy to Kubernetes
Update the `image` field in `kubernetes/deployment.yaml` with the Docker image URL:
```
<your-dockerhub-username>/<your-app-name>:<tag>
```

Apply the Kubernetes deployment and service configuration:
```bash
kubectl apply -f kubernetes/deployment.yaml
kubectl apply -f kubernetes/service.yaml
```

### 5. Get the External IP of the LoadBalancer
Retrieve the external IP address of the LoadBalancer by running:
```bash
kubectl get service/serversvc |  awk {'print $1" " $2 " " $4 " " $5'} | column -t
```
Look for the `EXTERNAL-IP` column in the output corresponding to your service. This is the public IP address you can use to access your application.

---

## Recap of Commands

### Local Deployment
1. **Docker Compose**
   ```bash
   docker-compose up --build
   ```

2. **kind**
   ```bash
   kind create cluster
   kubectl apply -f kubernetes/deployment.yaml
   kubectl port-forward service/<service-name> <local-port>:<service-port>
   ```
   **To update the image:**
   ```bash
   docker build -t <your-dockerhub-username>/<your-app-name>:<new-tag> .
   kubectl apply -f kubernetes/deployment.yaml
   ```

### Cloud Deployment
1. **Terraform**
   ```bash
   cd terraform
   terraform init
   terraform apply
   terraform destroy
   ```

2. **Docker**
   ```bash
   docker build -t <your-dockerhub-username>/<your-app-name>:<tag> .
   docker login
   docker push <your-dockerhub-username>/<your-app-name>:<tag>
   ```

3. **Kubernetes**
   ```bash
   kubectl apply -f kubernetes/deployment.yaml
   kubectl apply -f kubernetes/service.yaml
   kubectl get services
   ```

4. **Pod Access**
   ```bash
   kubectl get pods
   kubectl exec -it <pod-name> -- /bin/bash
   ```

---

## Notes
- Ensure you have proper permissions to manage AWS resources and Docker Hub repositories.
- Monitor your Kubernetes cluster using the `kubectl get pods` and `kubectl get services` commands to ensure everything is running smoothly.
- For local testing, kind is a lightweight and efficient way to simulate Kubernetes environments.
