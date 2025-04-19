#!/bin/bash

mkdir -p /output/agent &&
protoc --proto_path=/input \
    --go_out=/output/agent\
    --go_opt=paths=source_relative\
    --go-grpc_out=/output/agent \
    --go-grpc_opt=paths=source_relative \
    /input/agent.proto &&
chmod -R a+rw /output/agent &&
echo 'gRPC code generation complete. Files available in pkg/agent directory.'