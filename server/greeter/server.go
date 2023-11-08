package greeter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jafossum/grpc-web-streaming/greeter/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type greeterService struct {
	api.UnimplementedGreeterServer
}

func New() *greeterService {
	return &greeterService{}
}

func (s *greeterService) SayHello(ctx context.Context, hello *api.HelloRequest) (*api.HelloReply, error) {
	log.Println("Say Hello received")
	rep := &api.HelloReply{
		Message: hello.GetName(),
	}
	return rep, nil
}

func (s *greeterService) SayRepeatedHello(rect *api.RepeatHelloRequest, stream api.Greeter_SayRepeatedHelloServer) error {
	log.Println("Say Repeated Hello received")
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	i := 0
	for {
		i++
		if i > int(rect.GetCount()) {
			break
		}
		select {
		case <-tick.C:
			err := stream.Send(&api.HelloReply{
				Message: fmt.Sprintf("Name: %s, Count: %v", rect.GetName(), i),
			})
			if err != nil {
				return status.Error(codes.Aborted, "Issues!")
			}
		case <-stream.Context().Done():
			return nil
		}
	}
	return nil
}
