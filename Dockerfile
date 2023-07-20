FROM golang:1.20.4-alpine

WORKDIR /app

RUN apk upgrade --update && apk --no-cache add git

# COPY ./app/go .

ENV PATH=$PATH:/go/bin:$HOME/go/bin

RUN go install github.com/cosmtrek/air@latest

COPY ./app/go/go.mod ./app/go/go.sum ./
RUN go mod download

ENV DOCKERIZE_VERSION v0.6.1
RUN apk add --no-cache openssl \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

ENTRYPOINT dockerize -timeout 10s -wait tcp://db:3306  air -c .air.toml