#!/usr/bin/env bash
set -e

BUILD_TIME=$(date -Iseconds)
GO_VERSION=$(go version)
GIT_COMMIT=$(git rev-list -1 HEAD)
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
TARGET_DIR=${SND_RELEASE_DIR:='build/release/'}
APP_NAME=${SND_APP_NAME:='Sales & Dungeons'}

echo "Build Time : ${BUILD_TIME}"
echo "Branch     : ${GIT_BRANCH}"
echo "Commit     : ${GIT_COMMIT}"
echo "Go Version : ${GO_VERSION}"
echo "Build OS   : ${GOOS:=$(go env GOOS)}"
echo "Build Arch : ${GOARCH:=$(go env GOARCH)}"
echo "Build Tags : ${SND_TAGS}"

echo "Clearing old data..."
rm -r ${TARGET_DIR} || true

cd cmd/app
echo "Generating bundler config..."
JQ_CMD="""
.build_flags.tags = \"${SND_TAGS}\" |
.output_path = \"../../${TARGET_DIR}\" |
.environments[0].arch = \"${GOARCH}\" |
.environments[0].os = \"${GOOS}\"
"""
jq "${JQ_CMD}" -c bundler.json > bundler.gen.json

echo "Building App..."
astilectron-bundler -c bundler.gen.json -ldflags "X:github.com/BigJk/snd.GitCommitHash=${GIT_COMMIT}" -ldflags "X:github.com/BigJk/snd.GitBranch=${GIT_BRANCH}" -ldflags "X:github.com/BigJk/snd.BuildTime=${BUILD_TIME}"
cd ../..

# TODO: Rename output to match requested APP_NAME. app_name in bundler breaks if we pass anything with

echo "Building version.txt..."
echo "Commit: ${GIT_COMMIT}" > ${TARGET_DIR}/version.txt
echo "Branch: ${GIT_BRANCH}" >> ${TARGET_DIR}/version.txt
echo "Build Time: ${BUILD_TIME}" >> ${TARGET_DIR}/version.txt
