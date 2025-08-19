# Build stage
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

RUN go mod download 

# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server ./cmd

# Runtime stage
FROM gcr.io/distroless/static:nonroot 

WORKDIR /app
COPY --from=builder /app/server .

# Expose HTTP server
EXPOSE 8080

CMD ["./server"]