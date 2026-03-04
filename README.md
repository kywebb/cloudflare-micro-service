# Cloudflare Enterprise Audit Service

A high-performance, containerized Go microservice designed for **Edge-ready** environments. This service simulates a secure Enterprise RBAC (Role-Based Access Control) gateway with integrated audit logging.

## Key Features
* **Edge-Optimized:** Built with a Multi-Stage Dockerfile for a minimal (~15MB) footprint.
* **Enterprise RBAC:** Implements header-based authorization checks.
* **Structured Logging:** Outputs JSON audit logs to `stdout` for seamless integration with Cloudflare Logpush or SIEM tools.
* **Security First:** Runs as a non-privileged `appuser` within the container.

---

## Tech Stack
* **Language:** Go 1.23
* **Runtime:** Docker (Linux Alpine)
* **Pattern:** Twelve-Factor App (Stateless & Environment-driven)

---

## Quick Start

### 1. Build the Image
```bash
docker build -t cloudflare-micro-service:latest .
```

### 2. Run the Container
```bash
docker run -p 8080:8080 cloudflare-micro-service:latest
```

### 3. Test Admin Access (Success)
```bash
curl -i -H "X-User-Role: admin" http://localhost:8080/api/v1/access
```

### 4. Test Member Access (Forbidden)
```bash
curl -i -H "X-User-Role: member" http://localhost:8080/api/v1/access
```
