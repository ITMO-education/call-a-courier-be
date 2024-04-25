FROM golang as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /deploy/server/service ./cmd/service/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /deploy/server/ .
COPY --from=builder /app/config/ ./config/

LABEL MATRESHKA_CONFIG_ENABLED=true

EXPOSE 8080

RUN mkdir data

VOLUME /app/data

ENTRYPOINT ["./service"]