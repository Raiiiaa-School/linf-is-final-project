FROM golang:1.24 AS builder

WORKDIR /app

COPY pkg/go.mod pkg/go.sum pkg/ ./pkg/
COPY services/graphql/go.mod services/graphql/go.sum ./services/graphql/
COPY services/graphql/main.go ./services/graphql/
COPY services/graphql/graph/ ./services/graphql/graph/
COPY services/graphql/schema/ ./services/graphql/schema/

WORKDIR /app/services/graphql

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /graphql-service main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /graphql-service .

EXPOSE 8082

CMD [ "./graphql-service" ]