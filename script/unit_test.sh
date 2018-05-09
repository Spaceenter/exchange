#!/bin/sh

source env.sh
cd ${PROJECT_ROOT}
go test matching_engine/matcher/*.go
go test matching_engine/service/*.go
go test store/*.go
