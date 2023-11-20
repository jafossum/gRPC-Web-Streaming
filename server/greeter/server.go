package greeter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jafossum/grpc-web-streaming/greeter/api"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Subscriber interface {
	Subscribe(topic string) (<-chan *nats.Msg, error)
	Close()
}

type greeterService struct {
	api.UnimplementedGreeterServer
	client Subscriber
}

func New(sub Subscriber) (*greeterService, error) {
	return &greeterService{
		client: sub,
	}, nil
}

func (s *greeterService) Close() {
	s.client.Close()
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
			log.Println("Stream abandoned")
			return nil
		}
	}
	return nil
}

func (s *greeterService) SubscribeRepeatedHello(rect *api.SubscribeHelloRequest, stream api.Greeter_SubscribeRepeatedHelloServer) error {
	log.Println("Subscribe Repeated Hello received")

	subTs, err := s.client.Subscribe("foo.bar.timeseries")
	if err != nil {
		return status.Error(codes.Aborted, err.Error())
	}
	subHb, err := s.client.Subscribe("foo.bar.heartbeat")
	if err != nil {
		return status.Error(codes.Aborted, err.Error())
	}

	for {
		select {
		case msg := <-subTs:
			if msg == nil {
				continue
			}
			err := stream.Send(&api.HelloReply{
				Message: fmt.Sprintf("NATS -%s- Stream Receive: %s", msg.Subject, string(msg.Data)),
			})
			if err != nil {
				return status.Error(codes.Aborted, "Issues!")
			}
		case msg := <-subHb:
			if msg == nil {
				continue
			}
			err := stream.Send(&api.HelloReply{
				Message: fmt.Sprintf("NATS -%s- Stream Receive: %s", msg.Subject, string(msg.Data)),
			})
			if err != nil {
				return status.Error(codes.Aborted, "Issues!")
			}
		case <-stream.Context().Done():
			log.Println("Stream abandoned")
			return nil
		}
	}
}
