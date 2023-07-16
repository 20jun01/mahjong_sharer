FROM golang:1.20.4-alpine

WORKDIR  /app/go

# alpineパッケージのアップデート
RUN apk upgrade --update && apk --no-cache add git
RUN go install github.com/cosmtrek/air@latest

# ローカルの現在のディレクトリから、コンテナの作業ディレクトリにコピー
COPY . .

CMD ["air", "-c", ".air.toml"]