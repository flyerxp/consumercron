#!/usr/bin/env bash
RUN_NAME="consumer"
git checkout go.mod
git checkout ../crontab/go.mod
git pull
go mod tidy
mkdir -p ../output/bin ../output/conf
cp -r conf/* ../output/conf
go build -o ../output/bin/${RUN_NAME}
supervisorctl status all | grep consumer_ | awk '{print $1}' | xargs supervisorctl restart