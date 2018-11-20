#!/usr/bin/env bash

go test ./...  -run=parser_test.go -bench=. -benchmem


go test ./... -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1 parser_test.go