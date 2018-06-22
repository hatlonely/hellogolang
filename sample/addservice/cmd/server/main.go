package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/hashicorp/consul/api"
	addservice "github.com/hatlonely/hellogolang/sample/addservice/api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
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

// HealthImpl implement health
type HealthImpl struct{}

// Check service
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// ServiceRegister register service
func ServiceRegister(port int) {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	agent := client.Agent()
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("add-%v", port),
		Name:    "grpc.health.v1.addservice",
		Tags:    []string{"hatlonely"},
		Port:    port,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			Interval: "10s",
			GRPC:     fmt.Sprintf("127.0.0.1:%d/%s", port, "addservice"),
			DeregisterCriticalServiceAfter: "1m",
		},
	}
	if err := agent.ServiceRegister(reg); err != nil {
		panic(err)
	}
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
	grpc_health_v1.RegisterHealthServer(server, &HealthImpl{})

	port, _ := strconv.Atoi(os.Args[1])
	ServiceRegister(port)

	address, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Args[1]))
	if err != nil {
		panic(err)
	}

	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
