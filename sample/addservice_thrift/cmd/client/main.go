package main

import (
	"context"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/hatlonely/hellogolang/sample/addservice_thrift/api/addservice/gen-go/addservice"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("localhost:3001")
	if err != nil {
		panic(err)
	}
	transport, err = thrift.NewTBufferedTransportFactory(8192).GetTransport(transport)
	if err != nil {
		panic(err)
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		panic(err)
	}

	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := addservice.NewAddServiceClient(thrift.NewTStandardClient(iprot, oprot))

	limiter := rate.NewLimiter(rate.Every(time.Duration(800)*time.Millisecond), 1)

	hystrix.ConfigureCommand(
		"addservice",
		hystrix.CommandConfig{
			Timeout:                100,
			MaxConcurrentRequests:  2,
			RequestVolumeThreshold: 4,
			ErrorPercentThreshold:  25,
			SleepWindow:            1000,
		},
	)

	for a := int64(0); a < 10; a++ {
		for b := int64(0); b < 10; b++ {
			if err := limiter.Wait(context.Background()); err != nil {
				panic(err)
			}

			{
				var res *addservice.AddResponse
				req := &addservice.AddRequest{A: a, B: b}
				err := hystrix.Do("addservice", func() error {
					var err error
					ctx, cancel := context.WithTimeout(context.Background(), time.Duration(50*time.Millisecond))
					defer cancel()
					res, err = client.Add(ctx, req)
					return err
				}, func(err error) error {
					logrus.WithField("err", err).Error()
					res = &addservice.AddResponse{V: req.A + req.B}
					return nil
				})
				if err != nil {
					logrus.WithField("err", err).Error()
				}
				logrus.WithField("req", req).WithField("res", res).Info()
			}
		}
	}
}
