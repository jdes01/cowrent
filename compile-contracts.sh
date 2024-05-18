#!/bin/bash

set -x

CONTRACTS_DIR="contracts"
GO_OUT_DIR="api"
TS_OUT_DIR="web/contracts"
PROTOC_GEN_TS_PROTO="./web/node_modules/.bin/protoc-gen-ts_proto"

mkdir -p ${GO_OUT_DIR}
if [ $? -ne 0 ]; then
  echo "Error creating directory ${GO_OUT_DIR} for Go compilation"
  exit 1
fi

mkdir -p ${TS_OUT_DIR}
if [ $? -ne 0 ]; then
  echo "Error creating directory ${TS_OUT_DIR} for Typescript compilation"
  exit 1
fi

PROTO_FILES=$(find ${CONTRACTS_DIR} -name "*.proto")

for PROTO_FILE in $PROTO_FILES; do
  RELATIVE_PATH=$(realpath --relative-to=${CONTRACTS_DIR} ${PROTO_FILE})

  protoc --go_out=${GO_OUT_DIR} --go_opt=paths=source_relative ${PROTO_FILE}
  if [ $? -ne 0 ]; then
    echo "Error compiling ${PROTO_FILE} for Go"
    exit 1
  fi

  protoc --plugin=${PROTOC_GEN_TS_PROTO} --ts_proto_out=${TS_OUT_DIR} --proto_path=${CONTRACTS_DIR} ${RELATIVE_PATH}
  if [ $? -ne 0 ]; then
    echo "Error compiling ${PROTO_FILE} for TypeScript"
    exit 1
  fi
done

echo "Compilation completed successfully."
