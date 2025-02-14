#!/bin/bash

# docker-composeでコンテナを起動
docker-compose up -d

# コンテナイメージ一覧
docker image ls

# 実行中のコンテナ一覧
docker ps