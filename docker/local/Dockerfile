FROM golang:1.18-alpine AS builder

WORKDIR /lyfebloc-core

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o lyfebloc-core main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /lyfebloc-core

COPY --from=builder /lyfebloc-core/lyfebloc-core ./
COPY ./docker/local/lyfebloc-core.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478 5001

ENTRYPOINT ["./lyfebloc-core.sh"]
