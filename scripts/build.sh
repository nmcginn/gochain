#!/bin/bash -e
protoc -I=data/ --go_out=. data/test.proto
go fmt
go build
