FROM golang:1.23-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

FROM alpine:3.20

RUN addgroup -S app && adduser -S app -G app

USER app
WORKDIR /app

COPY --from=build /bin/api /app/api

EXPOSE 8080

ENTRYPOINT ["/app/api"]
