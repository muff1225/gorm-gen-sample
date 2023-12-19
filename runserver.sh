#!/bin/bash
base=${0##*/}

# docker-composeのaliasを作成
dc='docker-compose --env-file ./.env --file ./docker-compose.yml --project-name gorm-gen-sample'

# コンテナのビルドと起動
${dc} up -d --build

# appコンテナにログイン
${dc} exec app bash

# appコンテナから抜けたら自動的に停止とする
${dc} down --remove-orphans