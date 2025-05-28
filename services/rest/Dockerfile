FROM golang:1.24 AS builder

WORKDIR /app

COPY pkg/go.mod pkg/go.sum pkg/ ./pkg/
COPY services/rest/go.mod services/rest/go.sum ./services/rest/
COPY services/rest/main.go ./services/rest/
COPY services/rest/handlers/ ./services/rest/handlers

WORKDIR /app/services/rest

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /rest-service main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /rest-service .

EXPOSE 8081

CMD [ "./rest-service" ]