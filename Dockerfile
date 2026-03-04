FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]

# --- Stage 1: Build ---
FROM golang:1.23-alpine AS builder

# Add git/ca-certificates if your app needs to fetch private modules
RUN apk add --no-cache git ca-certificates

WORKDIR /app
COPY go.mod ./
# RUN go mod download # Uncomment if you have a go.mod file

COPY . .
# Build a static binary (important for scratch/alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o /enterprise-app .

# --- Stage 2: Final (The "Edge" Ready Image) ---
FROM alpine:latest

# Security: Don't run as root!
RUN adduser -D -u 1000 appuser
USER appuser

WORKDIR / 
COPY --from=builder /enterprise-app /enterprise-app

# Documentation Labels (OCI standards)
# LABEL org.opencontainers.image.source="https://github.com/yourusername/audit-service"
LABEL org.opencontainers.image.description="Go Audit Service for Enterprise RBAC"

EXPOSE 8080
ENTRYPOINT ["/enterprise-app"]

# docker run -p 8080:8080 audit-service:latest