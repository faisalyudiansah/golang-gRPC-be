#!/bin/bash

PROTO_DIR="./internal/grpc/proto"
OUT_DIR="./internal/grpc/proto/generate"

if ! command -v protoc &> /dev/null
then
    echo "Error: protoc tidak ditemukan. Silakan install protoc terlebih dahulu."
    exit 1
fi

if [ -d "$PROTO_DIR" ]; then
    echo "Generating gRPC code for $(basename $(pwd))..."
    rm -rf ${OUT_DIR}
    mkdir -p ${OUT_DIR}
    protoc \
        --proto_path=${PROTO_DIR} \
        --go_out=${OUT_DIR} \
        --go-grpc_out=${OUT_DIR} \
        --go_opt=paths=source_relative \
        --go-grpc_opt=paths=source_relative \
        ${PROTO_DIR}/*.proto
    echo "✅ gRPC code generated successfully!"
else
    echo "⚠️  No proto directory found, skipping gRPC code generation."
fi
