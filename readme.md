# Kubernetes Application Deployment Guide

## Prerequisites
Before getting started, ensure you have the following tools installed on your system:
- [Docker](https://www.docker.com)
- [AWS CLI](https://aws.amazon.com/cli/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

### Local Development Tools
For local development and testing, we recommend using:
- [Docker Compose](https://docs.docker.com/compose/)
- [kind](https://kind.sigs.k8s.io/) (Kubernetes in Docker)

## Local Deployment

### Using Docker Compose
To deploy the application locally using Docker Compose:

```bash
docker-compose up --build
```

This command builds and starts all services defined in your `docker-compose.yml` file.

### Using kind
1. Create a Kubernetes cluster:
   ```bash
   kind create cluster
   ```

2. Apply the Kubernetes deployment configuration:
   ```bash
   kubectl apply -f kubernetes/deployment.yaml
   ```

3. Access the application using port forwarding:
   ```bash
   kubectl port-forward service/<service-name> <local-port>:<service-port>
   ```

To update the Docker image in your kind cluster:
```bash
docker build -t <your-registry>/<image-name>:<tag> . && \
kubectl apply -f kubernetes/deployment.yaml
```

## Cloud Deployment

### Terraform Setup
1. Initialize the Terraform configuration:
   ```bash
   cd terraform && terraform init
   ```

2. Apply the infrastructure changes:
   ```bash
   terraform apply
   ```

3. Destroy the infrastructure when done:
   ```bash
   terraform destroy
   ```

### Docker Image Deployment
1. Build and tag your Docker image:
   ```bash
   docker build -t <your-registry>/<image-name>:<tag> .
   ```

2. Log in to Docker Hub:
   ```bash
   docker login
   ```

3. Push the image to Docker Hub:
   ```bash
   docker push <your-registry>/<image-name>:<tag>
   ```

### Kubernetes Deployment
1. Update your `kubernetes/deployment.yaml` file with the correct image URL.

2. Apply the deployment and service configurations:
   ```bash
   kubectl apply -f kubernetes/deployment.yaml
   kubectl apply -f kubernetes/service.yaml
   ```

3. Retrieve the external IP of your service:
   ```bash
   kubectl get services
   ```
   Look for the `EXTERNAL-IP` column to find your application's public endpoint.

## Monitoring and Maintenance

### Check Pod Status
```bash
kubectl get pods
```

### Access Running Pods
```bash
kubectl exec -it <pod-name> -- /bin/bash
```

### View Services
```bash
kubectl get services
```

## Notes
- Ensure you have the necessary permissions for AWS resources and Docker Hub repositories.
- Monitor your Kubernetes cluster using `kubectl get pods` and `kubectl get services`.
- For local testing, `kind` provides an efficient way to simulate Kubernetes environments.