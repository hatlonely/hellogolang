package main

import (
	"context"
	"math/rand"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/hatlonely/hellogolang/sample/addservice_thrift/api/addservice/gen-go/addservice"
	"github.com/sirupsen/logrus"
)

// AddServiceImpl 实现 Add 服务
type AddServiceImpl struct{}

// Add 接口实现
func (s *AddServiceImpl) Add(ctx context.Context, request *addservice.AddRequest) (*addservice.AddResponse, error) {
	// 50% 概率 sleep，模拟超时场景
	if rand.Int()%2 == 0 {
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	response := &addservice.AddResponse{
		V: request.A + request.B,
	}
	logrus.WithField("request", request).WithField("response", response).Info()
	return response, nil
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	transport, err := thrift.NewTServerSocket(":3001")
	if err != nil {
		panic(err)
	}

	processor := addservice.NewAddServiceProcessor(&AddServiceImpl{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
	)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
