#!/bin/bash -e
protoc -I=data/ --go_out=. data/*.proto
go fmt
go build
