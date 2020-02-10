FROM golang:1.13.0-alpine

ENV USER DEV
ENV GOPATH /go:$GOPATH
ENV PATH /go/bin:$PATH

ENV SERVICE_NAME=minilru \
    SERVICE_DESC='MiniLru' \
    SERVICE_TAGS='local,local-test,minilru-cloud,minilru' \
    SERVICE_CHECK_ENDPOINT='/' \
    SERVICE_CHECK_INTERVAL=10s \
    SERVICE_CHECK_TIMEOUT=2s \
    SERVICE_DB_CONFIG='[{"type":"postgresql","postgresql_list":[{"server":"database-postgres-dev", "port": "5432", "databases":[{"name":"dbpostgres","username":"postgres","password":"postgresqL"}]}]}]' \
    SERVICE_CACHE_CONFIG='[{"type":"redis","redis_list":[{"server":"redis-dev", "port": "6379", "databases":[{"name":"0","username":"","password":""}]}]}]' \
    SERVICE_SALT='SaltNPepa' \
    NODE_ENV=dev \
    PORT=8081

RUN apk add --update --no-cache alpine-sdk bash ca-certificates \
      libressl \
      tar \
      git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils

RUN mkdir /app
RUN apk update && apk add --no-cache git && apk add --no-cache openssh
ADD ./app /app
WORKDIR /app/src