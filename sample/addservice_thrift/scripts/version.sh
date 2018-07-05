#!/usr/bin/env bash

export PATH=$(pwd)/third/go/bin:$PATH

echo "appversion: $(git describe --tags) [git describe --tags]
gitremote: $(git remote -v | grep fetch | awk '{print $2}') [git remote -v | grep fetch]
hashcode: $(git rev-parse HEAD) [git rev-parse HEAD]
datetime: $(date '+%Y-%m-%d %H:%M:%S') [date]
hostname: $(hostname):$(pwd) [hostname:pwd]
goversion: $(go version) [go version]"
