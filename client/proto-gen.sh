protoc -I=. -I=.. --js_out=import_style=commonjs:. ../greeter-service.proto
protoc -I=. -I=.. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. ../greeter-service.proto
