#!/bin/bash

set -e

PROJECT_PATH="$GOPATH/src/github.com/$(git config user.name)/GoProjects"
ORG_NAME="DESMG"
PROJECT_NAME="StatusAgent"
EXECUTABLE_NAME="StatusAgent"

TARGET_ARCHS=(
    "linux/amd64"
    "windows/amd64"
    "darwin/arm64"
)

for arch in "${TARGET_ARCHS[@]}"; do
    os=$(echo "$arch" | cut -d'/' -f1)
    arch=$(echo "$arch" | cut -d'/' -f2)
    echo "Building for $os/$arch..."
    GOOS=$os GOARCH=$arch go build -o "$PROJECT_PATH/dist/$os-$arch/$EXECUTABLE_NAME" "$PROJECT_PATH/$ORG_NAME-$PROJECT_NAME" &
done

wait
