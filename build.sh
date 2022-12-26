#!/usr/bin/env bash
set -e

BUILD_TIME=$(date -Iseconds)
GO_VERSION=$(go version)
GIT_COMMIT=$(git rev-list -1 HEAD)
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

echo "Build Time : ${BUILD_TIME}"
echo "Branch     : ${GIT_BRANCH}"
echo "Commit     : ${GIT_COMMIT}"
echo "Go Version : ${GO_VERSION}"
echo "Build OS   : ${GOOS:=$(go env GOOS)}"
echo "Build Arch : ${GOARCH:=$(go env GOARCH)}"
echo "Build Tags : ${SND_TAGS}"

echo "Clearing old data..."
rm -r build || true
mkdir -p build

echo "Building App..."
case "${GOOS}" in
  "windows") EXT=".exe" ;;
  *) EXT="" ;;
esac
LD_FLAGS="-X github.com/BigJk/snd.GitCommitHash=${GIT_COMMIT} -X github.com/BigJk/snd.GitBranch=${GIT_BRANCH} -X github.com/BigJk/snd.BuildTime=${BUILD_TIME}"

cd cmd/app
go build -ldflags "${LD_FLAGS}" -o app -tags "${SND_TAGS}"
cd ../..
mv cmd/app/app "build/SND${EXT}"

echo "Building version.txt..."
echo "Commit: ${GIT_COMMIT}" > build/version.txt
echo "Branch: ${GIT_BRANCH}" >> build/version.txt
echo "Build Time: ${BUILD_TIME}" >> build/version.txt
