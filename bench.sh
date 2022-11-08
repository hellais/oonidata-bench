#!/bin/sh
go get ./...
go build bench_go.go
go build bench_go_goccy.go
poetry run python bench_python.py && ./bench_go && ./bench_go_goccy
