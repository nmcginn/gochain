#!/bin/bash -ex
protoc -I=data/ --go_out=. data/*.proto
go fmt
go build
go test
