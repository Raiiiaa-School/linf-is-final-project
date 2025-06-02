FROM golang:1.24 AS builder

WORKDIR /app

COPY pkg/go.mod pkg/go.sum pkg/ ./pkg/
COPY services/soap/go.mod services/soap/go.sum ./services/soap/
COPY services/soap/main.go ./services/soap/
COPY services/soap/handlers/ ./services/soap/handlers/

WORKDIR /app/services/soap

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /soap-service main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /soap-service .

EXPOSE 8084

CMD [ "./soap-service" ]