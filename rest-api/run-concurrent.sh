#!/usr/bin/env bash
echo Endpoint is available at http://127.0.0.1:8082/concurrent
echo Run benchmark: ab -n 100 -c 1 http://127.0.0.1:8082/concurrent
go run concurrent/main.go | logstalgia --simulation-speed 5 --sync -
