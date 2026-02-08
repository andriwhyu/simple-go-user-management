FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o user-management ./cmd/api

# -------- Runtime stage --------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/user-management /app/user-management

USER nonroot:nonroot

ENTRYPOINT ["/app/user-management"]