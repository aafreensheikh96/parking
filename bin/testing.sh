#!/bin/sh
go build ./cmd/parking/main.go
go run ./cmd/parking/main.go ./sample/input.txt
