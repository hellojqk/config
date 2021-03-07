#! /bin/bash

# cp -r $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.1.0/third_party/googleapis/google ./google

protoc -I . \
    --go_out ./gen/go/ --go_opt paths=source_relative \
    --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./gen/go \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true \
    proto/*.proto
