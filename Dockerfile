FROM golang:1.13-alpine

LABEL maintainer="bondhan novandy<bondhan.novandy@gmail.com>"

ENV TZ=Asia/Jakarta
ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN cd ./cmd/gotem/ && go mod download && go build -o /app/gotem main.go

RUN apk add --no-cache --virtual .build-deps --no-progress -q \
    curl \
    tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 8080

CMD ["./gotem"]
