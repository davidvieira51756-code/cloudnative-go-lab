# cloudnative-go-lab

Small hands-on project to learn Go, Docker, Kubernetes, and cloud-native fundamentals.

## Features

- Minimal HTTP API using the Go standard library
- `GET /`
- `GET /health`
- Configurable `PORT`
- Graceful shutdown
- Basic tests
- Multi-stage Docker build
- Kubernetes Deployment, Service, and ConfigMap
- Liveness and readiness probes
- GitHub Actions CI

## Run locally

```bash
go run ./cmd/api
curl http://localhost:8080/health
```

## Docker

```bash
docker build -t cloudnative-go-lab .
docker run --rm -p 8080:8080 cloudnative-go-lab
```

## Running locally with Kubernetes (kind)

Create a local Kubernetes cluster with `kind`:

```bash
kind create cluster --name cloudnative-go-lab
```

Build the Docker image locally:

```bash
docker build -t cloudnative-go-lab:local .
```

Load the local image into the `kind` cluster:

```bash
kind load docker-image cloudnative-go-lab:local --name cloudnative-go-lab
```

Apply the Kubernetes manifests:

```bash
kubectl apply -f k8s/
```

Verify the Deployment, Pods, and Service:

```bash
kubectl get deployments
kubectl get pods
kubectl get services
kubectl describe deployment cloudnative-go-lab
```

Forward the Service to your local machine:

```bash
kubectl port-forward svc/cloudnative-go-lab 8080:80
```

In another terminal, test the API:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/
```

Destroy the cluster when you are done:

```bash
kind delete cluster --name cloudnative-go-lab
```

## CI

GitHub Actions runs on push and pull requests:

- `gofmt`
- `go vet`
- `go test`
- `go build`

## Next steps

- Metrics and observability
- Secrets and configuration management
- Rolling updates
- Autoscaling
- Kubernetes Operators in Go
