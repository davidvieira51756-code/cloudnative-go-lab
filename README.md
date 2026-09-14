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

## Kubernetes

```bash
kubectl apply -f k8s/
kubectl get pods
kubectl get svc
kubectl port-forward svc/cloudnative-go-lab 8080:80
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