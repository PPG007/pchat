#!/usr/bin/zsh

OUT_DIR="./src/pb"
SRC_DIR="../proto"

rm -rf ${OUT_DIR}/**/*.ts

protoc \
  --plugin=./node_modules/.bin/protoc-gen-ts_proto \
  --ts_proto_out=${OUT_DIR} \
  --ts_proto_opt=esModuleInterop=true \
  --ts_proto_opt=importSuffix=.js \
  --proto_path="${SRC_DIR}" \
  -I ${SRC_DIR} \
  ${SRC_DIR}/**/*.proto

../backend/scripts/error/gen
../backend/scripts/permission/gen
