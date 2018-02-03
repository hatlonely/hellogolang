package main

import (
	"hellogolang/api/counter_proto"
	"fmt"
	"time"
	"net"
	"google.golang.org/grpc"
)

type CounterServerImp struct {

}

func (c *CounterServerImp) Count(req *counter.CountReq, stream counter.Counter_CountServer) error {
	fmt.Printf("request from client. start: [%v]\n", req.GetStart())

	i := req.GetStart()
	for {
		i++
		stream.Send(&counter.CountRes{
			Num: i,
		})
		time.Sleep(time.Duration(500) * time.Millisecond)
	}

	return nil
}

func main() {
	server := grpc.NewServer()
	counter.RegisterCounterServer(server, &CounterServerImp{})

	address, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
