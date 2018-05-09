#!/bin/sh

source env.sh
cd ${GOPATH}/src
protoc --proto_path=${GOPATH}/src --go_out=${GOPATH}/src ${REL_PROJECT_ROOT}/matching_engine/matcher/proto/*.proto
protoc --proto_path=${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${REL_PROJECT_ROOT}/matching_engine/rpc/proto/*.proto
