#!/bin/sh

rm -rf build
mkdir -p build

go build -o build/stock-market-bot ./src/*.go
