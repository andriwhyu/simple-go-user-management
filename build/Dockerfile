# syntax=docker/dockerfile:1.27

FROM --platform=$BUILDPLATFORM golang:1.23.4-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy source
COPY . .

ARG TARGETOS
ARG TARGETARCH

# Build a static binary
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build,id=gobuild-${TARGETOS}-${TARGETARCH} \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" \
    -o /out/user-management ./cmd/api

# -------- Runtime stage --------
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /out/user-management /app/user-management

ENTRYPOINT ["/app/user-management"]