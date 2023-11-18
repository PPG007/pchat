#!/usr/bin/zsh
OUT_DIR="./pb"
SRC_DIR="../proto"

rm -rf ${OUT_DIR}/*.go

protoc -I ${SRC_DIR} \
  --go_out=${OUT_DIR} \
  --go_opt paths=source_relative  ${SRC_DIR}/*.proto

protoc-go-inject-tag -input="${OUT_DIR}/*.pb.go"
