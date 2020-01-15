FROM golang:1.13-alpine AS build-env

LABEL maintainer="bondhan novandy<bondhan.novandy@gmail.com>"

ENV TZ=Asia/Jakarta
ENV GO111MODULE=on

RUN apk update && apk upgrade
RUN apk add --no-cache --virtual .build-deps --no-progress -q \
    bash \
    curl \
    busybox-extras \
    tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Install dep if needed
# ENV DEP_VERSION="0.4.1"
# RUN curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $GOPATH/bin/dep
# RUN chmod +x $GOPATH/bin/dep
# RUN dep ensure

COPY . $GOPATH/src/github.com/bondhan/gotem
WORKDIR $GOPATH/src/github.com/bondhan/gotem

RUN cd ./cmd/gotem/ && go mod download && go build -o /app/gotem main.go

FROM alpine

WORKDIR /app
COPY --from=build-env $GOPATH/src/github.com/bondhan/gotem /app/gotem

EXPOSE 1323

CMD ["./gotem"]
