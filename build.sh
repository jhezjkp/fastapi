#!/bin/bash

# check if fastapi file exists, the delete it
if [ -f "fastapi" ]; then
  echo "fastapi exists, deleting..."
  rm fastapi
fi

# build new fastapi
echo "building fastapi..."

BUILD_DATE="$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
GIT_COMMIT="$(git rev-parse HEAD)"
VERSION="$(git describe --tags --abbrev=0 | tr -d '\n')"

GOOS=linux GOARCH=amd64 go build -o fastapi -ldflags="-X 'github.com/iimeta/fastapi/internal/version.buildDate=${BUILD_DATE}' -X 'github.com/iimeta/fastapi/internal/version.gitCommit=${GIT_COMMIT}' -X 'github.com/iimeta/fastapi/internal/version.gitVersion=${VERSION}'" main.go

echo "fastapi build complete!"