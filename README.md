# gRPC-Web Streaming example

Combining information from these sources:

- [gRPC-Web Basic Tutorial](https://grpc.io/docs/platforms/web/basics/)
- [gRPC-Web GitHub Examples](https://github.com/grpc/grpc-web/tree/master/net/grpc/gateway/examples/helloworld)

## Genarate Code

### Generate Server and Client proto files

    protoc -I=. --go_out=./server --go-grpc_out=./server ./greeter-service.proto
    protoc -I=. --js_out=import_style=commonjs:./client ./greeter-service.proto
    protoc -I=. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client ./greeter-service.proto

### Generate client code

    cd client
    (May be needed) export NODE_OPTIONS=--openssl-legacy-provider
    npm install
    npx webpack client.js

## Start the system

### Server

    cd server
    go run main.go

### Envoy

    docker run --rm -d -v "$(pwd)"/envoy.yaml:/etc/envoy/envoy.yaml:ro -p 8080:8080 -p 9901:9901 envoyproxy/envoy:v1.22.0

### Run the client in a webserver

    cd client
    python3 -m http.server 8081 &

### Nice to know commands

    jobs

    fg
    fg %1

    bg
    bg %1

    kill %1

## See the Result

Visit `localhost:8081` and watch the console (Chrome: `F12`).
