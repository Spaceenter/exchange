#!/bin/sh

source env.sh
cd ${PROJECT_ROOT}

go test matching_engine/matcher/*.go
go test matching_engine/rpc/client/*.go
go test matching_engine/rpc/service/*.go
go test store/*.go
