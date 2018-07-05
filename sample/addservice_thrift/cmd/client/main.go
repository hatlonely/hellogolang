package main

import (
	"context"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/hatlonely/hellogolang/sample/addservice_thrift/api/addservice/gen-go/addservice"
)

func main() {
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

	var res *addservice.AddResponse
	res, err = client.Add(context.Background(), &addservice.AddRequest{
		A: 1,
		B: 2,
	})

	fmt.Println(res)
}
