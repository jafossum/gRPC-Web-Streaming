package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jafossum/grpc-web-streaming/greeter"
	"github.com/jafossum/grpc-web-streaming/greeter/api"
	"github.com/jafossum/grpc-web-streaming/nats"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("a", "localhost", "Host Adress")
	port := flag.Int("p", 9090, "Port")

	flag.Parse()

	listen := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("Listening on:", listen)

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	sub := nats.NewNats()
	grSrv, err := greeter.New(sub)
	if err != nil {
		log.Fatal(err)
	}
	api.RegisterGreeterServer(grpcServer, grSrv)

	grpcServer.Serve(lis)
}
