FROM golang:1.20.4-alpine

WORKDIR /app

RUN apk upgrade --update && apk --no-cache add git

COPY ./app/go .

ENV PATH=$PATH:/go/bin:$HOME/go/bin

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
