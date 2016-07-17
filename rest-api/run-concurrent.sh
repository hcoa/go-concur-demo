#!/usr/bin/env bash
echo Endpoint is available at http://127.0.0.1:8080/concurrent
echo Run benchmark: ab -n 100 -c 4 http://127.0.0.1:8080/concurrent
go run concurrent/main.go | logstalgia --sync -