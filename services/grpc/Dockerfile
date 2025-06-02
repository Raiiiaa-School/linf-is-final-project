FROM golang:1.24 AS builder

# Install protoc and required dependencies
RUN apt-get update && apt-get install -y protobuf-compiler

WORKDIR /app

# Copy go.mod files
COPY pkg/go.mod pkg/go.sum pkg/ ./pkg/
COPY services/grpc/go.mod services/grpc/go.sum ./services/grpc/

# Install protoc-gen-go and protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

# Copy source code
COPY services/grpc/main.go ./services/grpc/
COPY services/grpc/handlers/ ./services/grpc/handlers/
COPY services/grpc/proto/ ./services/grpc/proto/

WORKDIR /app/services/grpc

# Download dependencies
RUN go mod download

# Generate protobuf code
RUN protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    proto/service.proto

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /grpc-service main.go

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /grpc-service .

EXPOSE 8083

CMD [ "./grpc-service" ]
