#!/bin/bash

# check if fastapi file exists, the delete it
if [ -f "fastapi" ]; then
  echo "fastapi exists, deleting..."
  rm fastapi
fi

# build new fastapi
echo "building fastapi..."
GOOS=linux GOARCH=amd64 go build -o fastapi main.go
echo "fastapi build complete!"