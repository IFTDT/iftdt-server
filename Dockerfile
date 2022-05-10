FROM golang:1.18-alpine3.15 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /app

ADD . /app

RUN go build main.go

FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .

COPY wait-for.sh .

EXPOSE 8080

CMD ["./main"]