FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN apk update &&  apk add git
RUN go get github.com/cosmtrek/air@latest

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
CMD ["air", "-c", "./.air.toml"]