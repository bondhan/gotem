FROM golang:1.13-alpine AS build-env

LABEL maintainer="bondhan novandy<bondhan.novandy@gmail.com>"

ENV GO111MODULE=on
ENV PROJECT=src/github.com/bondhan/gotem

COPY . $GOPATH/$PROJECT
WORKDIR $GOPATH/$PROJECT

RUN cd ./cmd/gotem/ && go mod download && go build -o $GOPATH/$PROJECT/gotem main.go


# clean container
FROM alpine

ENV TZ=Asia/Jakarta

RUN apk update && apk upgrade
RUN apk add --no-cache --virtual .build-deps --no-progress -q \
    bash \
    curl \
    busybox-extras \
    tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=build-env /go/src/github.com/bondhan/gotem/gotem /app/gotem
COPY --from=build-env /go/src/github.com/bondhan/gotem/.env /app/.env
COPY --from=build-env /go/src/github.com/bondhan/gotem/wait-for-it.sh /app/wait-for-it.sh

# for mysql
# CMD ./wait-for-it.sh --host=mysql-svc --port=3306 --timeout=60 -- ./gotem

# for postgres
CMD ./wait-for-it.sh --host=postgres-svc --port=5432 --timeout=60 -- ./gotem

EXPOSE 1323

# CMD ["./gotem"]

