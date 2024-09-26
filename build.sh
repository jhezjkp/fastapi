#!/bin/bash

# check if fastapi file exists, the delete it
target_file="deploy/fastapi"
if [ -f "${target_file}" ]; then
  echo "${target_file} exists, deleting..."
  rm ${target_file}
fi

# build new fastapi
echo "building fastapi..."

BUILD_DATE="$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
GIT_COMMIT="$(git rev-parse HEAD)"
VERSION="$(git describe --tags --abbrev=0 | tr -d '\n')"

GOOS=linux GOARCH=amd64 go build -o ${target_file} -ldflags="-X 'github.com/iimeta/fastapi/internal/version.buildDate=${BUILD_DATE}' -X 'github.com/iimeta/fastapi/internal/version.gitCommit=${GIT_COMMIT}' -X 'github.com/iimeta/fastapi/internal/version.gitVersion=${VERSION}'" main.go

chmod +x ${target_file}

tar -zcvf ${target_file}.tar.gz ${target_file}

echo "fastapi build complete!"