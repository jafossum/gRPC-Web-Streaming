PROTOC_GEN_TS_PROTO_PATH="./node_modules/.bin/protoc-gen-ts_proto"

protoc -I=.. -I=. --plugin="${PROTOC_GEN_TS_PROTO_PATH}" --ts_proto_out=./src/app/greeter-client/proto --ts_proto_opt=outputClientImpl=grpc-web ../*.proto
