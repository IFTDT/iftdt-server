FROM golang:1.18

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /app

ADD . /app

RUN go build main.go

EXPOSE 8080

# ENTRYPOINT ["./main"]