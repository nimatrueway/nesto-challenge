FROM golang:1.22.4-alpine3.20 as base

WORKDIR /app

COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

COPY cmd/ ./cmd
COPY internal/ ./internal
RUN go build -o readcommend ./cmd/server

FROM alpine:3.20

ENV GIN_MODE=release

WORKDIR /app

COPY --from=base /app/readcommend /app/readcommend

EXPOSE 5001

CMD ["./readcommend"]