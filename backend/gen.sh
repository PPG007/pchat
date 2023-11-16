#!/usr/bin/zsh
rm -rf ./pb/*.go
protoc -I ../proto \
  --go_out=./pb \
  --go_opt paths=source_relative  ../proto/*.proto
