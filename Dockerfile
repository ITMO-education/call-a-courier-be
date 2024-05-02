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
        cp -r config /deploy/server/config && \
        [ -d "./migrations" ] && cp -r ./migrations /deploy/server/migrations
FROM alpine

WORKDIR /app

COPY --from=builder /deploy/server/ .

LABEL MATRESHKA_CONFIG_ENABLED=true

EXPOSE 8080

ENTRYPOINT ["./service"]