FROM --platform=$BUILDPLATFORM golang as builder

WORKDIR /app

ARG TARGETOS
ARG TARGETARCH

COPY . .

RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
        go build -o /deploy/server/service ./cmd/service/main.go && \
        cp -r config /deploy/server/config
FROM alpine

WORKDIR /app

COPY --from=builder /deploy/server/ .
COPY --from=builder /app/config/ ./config/

LABEL MATRESHKA_CONFIG_ENABLED=true

EXPOSE 8080

ENTRYPOINT ["./service"]