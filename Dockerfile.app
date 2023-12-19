# 軽量なalpineイメージを使う
FROM golang:1.21-alpine3.18 AS golang-build

# Goビルドに必要なパッケージ群をインストールする
RUN apk add --no-cache git make gcc musl-dev bash

# タイムゾーンをJSTに設定する
RUN apk --no-cache add tzdata && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

# 開発環境にはgolang-migrateを入れる
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 依存パッケージのインストール
WORKDIR /go/src/github.com/muff1225/gorm-gen-sample/app