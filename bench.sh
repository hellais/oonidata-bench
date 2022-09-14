#!/bin/sh

poetry run python bench_python.py && go run bench_go.go
