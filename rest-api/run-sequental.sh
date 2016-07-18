#!/usr/bin/env bash
echo Endpoint is available at http://127.0.0.1:8081/sequential
echo Run benchmark: ab -n 100 -c 1 http://127.0.0.1:8081/sequential
go run sequential/main.go | logstalgia --simulation-speed 5 --sync -
