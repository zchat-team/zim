#!/bin/bash

echo 'Generating protocol'
protoc -I. -I./third_party --gofast_out=. --gofast_opt paths=source_relative ./app/conn/protocol/protocol.proto

echo 'Generating rpc api'
PROTOS=$(find ./proto/rpc -type f -name '*.proto')

for PROTO in $PROTOS; do
  echo $PROTO
  protoc \
    -I. -I./proto/rpc/common -I$(dirname $PROTO) \
    -I./third_party \
    --gofast_out=. \
    --gofast_opt paths=source_relative \
    --rpcx_out=. \
    --rpcx_opt paths=source_relative \
    $PROTO
done

echo 'Generating http api'
PROTOS=$(find ./proto/http -type f -name '*.proto')

for PROTO in $PROTOS; do
  echo $PROTO
  protoc \
    -I. \
    -I$(dirname $PROTO) \
    -I./third_party \
    --gofast_out=. \
    --gofast_opt paths=source_relative \
    --zmicro-gin_out=. \
    --zmicro-gin_opt paths=source_relative \
    --zmicro-gin_opt allow_empty_patch_body=true \
    $PROTO
done

echo 'Generating errno'
ERRORS=$(find ./errno -type f -name '*.proto')
for ERROR in $ERRORS; do
  echo $ERROR
  protoc \
  -I. -I${GOPATH}/src \
  --gofast_out=. \
  --gofast_opt paths=source_relative \
  --zmicro-errno_out=. \
  --zmicro-errno_opt epk=github.com/zmicro-team/zmicro/core/errors \
  --zmicro-errno_opt paths=source_relative \
  $ERROR
done

echo 'Generating rest api swagger'
PROTOS=$(find ./proto/http/rest -type f -name '*.proto')
protoc \
    -I . \
    -I./third_party \
    --openapiv2_out ./app/rest/docs \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt allow_merge=true \
    --openapiv2_opt merge_file_name=swagger \
    --openapiv2_opt enums_as_ints=true \
    --openapiv2_opt json_names_for_fields=false \
    $PROTOS

echo 'Generating demo api swagger'
PROTOS=$(find ./proto/http/demo -type f -name '*.proto')
protoc \
    -I . \
    -I./third_party \
    --openapiv2_out ./app/demo/docs \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt allow_merge=true \
    --openapiv2_opt merge_file_name=swagger \
    --openapiv2_opt enums_as_ints=true \
    --openapiv2_opt json_names_for_fields=false \
    $PROTOS