#!/usr/bin/env bash
RUN_NAME="crontab"
git checkout go.mod
git pull
go mod tidy
go build -o ../output/bin/${RUN_NAME}
#supervisorctl status all | grep crontab_ | awk '{print $1}' | xargs supervisorctl restart