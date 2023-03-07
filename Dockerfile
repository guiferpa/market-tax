FROM golang:1.19 AS builder

WORKDIR /opt/app

COPY . .

RUN GCO_ENABLED=0 go build -o ./dist/cli ./cmd/cli/main.go

FROM scratch

COPY  --from=builder /opt/app/dist/cli /usr/local/bin/app

ENTRYPOINT ["app"]
