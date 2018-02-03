package main

import (
	"google.golang.org/grpc"
	"fmt"
	"golang.org/x/net/context"
	"os"
	"hellogolang/api/counter_proto"
	"strconv"
)

func main() {
	start, _ := strconv.ParseInt(os.Args[1], 10, 64)

	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("dial failed. err: [%v]\n", err)
		return
	}
	client := counter.NewCounterClient(conn)

	stream, err := client.Count(context.Background(), &counter.CountReq{
		Start: start,
	})
	if err != nil {
		fmt.Errorf("count failed. err: [%v]\n", err)
		return
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			fmt.Errorf("client count failed. err: [%v]", err)
			return
		}

		fmt.Printf("server count: %v\n", res.GetNum())
	}
}
