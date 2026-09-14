# cloudnative-go-lab

Projeto pessoal pequeno para aprender Go, Docker e Kubernetes de forma pratica, sem microservicos artificiais nem dependencias desnecessarias.

## Arquitetura

- `cmd/api`: ponto de entrada da API.
- `internal/http`: handlers HTTP e router da aplicacao.
- `k8s`: manifests Kubernetes para Deployment, Service e ConfigMap.
- `.github/workflows`: pipeline simples de CI.

A aplicacao usa apenas a standard library do Go. Expoe:

- `GET /`: mensagem simples sobre a aplicacao.
- `GET /health`: resposta JSON com `{"status":"ok"}`.

A porta e configurada pela variavel de ambiente `PORT` e usa `8080` por omissao.

## Correr localmente

```bash
go run ./cmd/api
```

Com porta personalizada:

```bash
PORT=9090 go run ./cmd/api
```

Em PowerShell:

```powershell
$env:PORT = "9090"
go run ./cmd/api
```

Testar os endpoints:

```bash
curl http://localhost:8080/
curl http://localhost:8080/health
```

Executar checks locais:

```bash
gofmt -w .
go vet ./...
go test ./...
go build ./...
```

## Correr com Docker

Construir a imagem:

```bash
docker build -t cloudnative-go-lab:latest .
```

Correr o container:

```bash
docker run --rm -p 8080:8080 -e PORT=8080 cloudnative-go-lab:latest
```

Testar:

```bash
curl http://localhost:8080/health
```

## Kubernetes

Aplicar os manifests:

```bash
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

Verificar recursos:

```bash
kubectl get pods
kubectl get svc cloudnative-go-lab
```

Para testar localmente com port-forward:

```bash
kubectl port-forward svc/cloudnative-go-lab 8080:80
curl http://localhost:8080/health
```

Nota: se estiveres a usar um cluster local como kind ou minikube, garante que a imagem `cloudnative-go-lab:latest` esta disponivel dentro do cluster.

## CI

O GitHub Actions corre em `push` e `pull_request`:

- `gofmt` check;
- `go vet ./...`;
- `go test ./...`;
- `go build ./...`.

## Proximos passos possiveis

- adicionar metrics;
- observability;
- ConfigMaps/Secrets;
- rolling updates;
- autoscaling;
- experimentar Kubernetes Operators em Go posteriormente.
