#!/bin/sh

CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/internal/pkg/genproto
mkdir -p ${CURRENT_DIR}/internal/pkg/genproto

for x in $(find ${CURRENT_DIR}/submodule* -type d); do
  protoc -I=${x} -I/usr/local/include \
    --go_out=${CURRENT_DIR}/internal/pkg/genproto --go_opt=paths=source_relative \
    --go-grpc_out=${CURRENT_DIR}/internal/pkg/genproto --go-grpc_opt=paths=source_relative ${x}/*.proto
done
