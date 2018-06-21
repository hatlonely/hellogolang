package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	addservice "github.com/hatlonely/hellogolang/sample/addservice/api"
)

// AddServiceImpl implement add services
type AddServiceImpl struct{}

// Add implemention
func (s *AddServiceImpl) Add(ctx context.Context, request *addservice.AddRequest) (*addservice.AddResponse, error) {
	// 50% 概率 sleep，模拟超时场景
	if rand.Int()%2 == 0 {
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	fmt.Println(request)
	return &addservice.AddResponse{
		V: request.A + request.B,
	}, nil
}

func main() {
	server := grpc.NewServer()
	addservice.RegisterAddServiceServer(server, &AddServiceImpl{})

	address, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
