package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	addservice "github.com/hatlonely/hellogolang/sample/addservice/api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// AddServiceImpl implement add services
type AddServiceImpl struct{}

// Add implemention
func (s *AddServiceImpl) Add(ctx context.Context, request *addservice.AddRequest) (*addservice.AddResponse, error) {
	// 50% 概率 sleep，模拟超时场景
	if rand.Int()%2 == 0 {
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	// fmt.Println(request)
	return &addservice.AddResponse{
		V: request.A + request.B,
	}, nil
}

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	entry := logrus.NewEntry(logger)
	grpc_logrus.ReplaceGrpcLogger(entry)

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
			),
			grpc_logrus.UnaryServerInterceptor(
				entry,
				grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
					return "grpc.time_ns", duration.Nanoseconds()
				}),
			),
			grpc_logrus.PayloadUnaryServerInterceptor(entry, func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true }),
		),
	)
	addservice.RegisterAddServiceServer(server, &AddServiceImpl{})

	address, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Args[1]))
	if err != nil {
		panic(err)
	}

	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
