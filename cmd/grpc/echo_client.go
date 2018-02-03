package main

import (
	"google.golang.org/grpc"
	"fmt"
	"hellogolang/api/echo_proto"
	"golang.org/x/net/context"
	"os"
	"strings"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("dial failed. err: [%v]\n", err)
		return
	}

	client := echo.NewEchoClient(conn)
	res, err := client.Echo(context.Background(), &echo.EchoReq{
		Msg: strings.Join(os.Args[1:], " "),
	})

	if err != nil {
		fmt.Errorf("client echo failed. err: [%v]", err)
		return
	}

	fmt.Printf("message from server: %v", res.GetMsg())
}
