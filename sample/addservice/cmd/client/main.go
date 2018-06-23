package main

import (
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	addservice "github.com/hatlonely/hellogolang/sample/addservice/api"
	"github.com/hatlonely/hellogolang/sample/addservice/internal/grpclb"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	conn, err := grpc.Dial(
		"",
		grpc.WithInsecure(),
		// 开启 grpc 中间件的重试功能
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Duration(1)*time.Millisecond)), // 重试间隔时间
				grpc_retry.WithMax(3),                                             // 重试次数
				grpc_retry.WithPerRetryTimeout(time.Duration(5)*time.Millisecond), // 重试时间
				// 返回码为如下值时重试
				grpc_retry.WithCodes(codes.ResourceExhausted, codes.Unavailable, codes.DeadlineExceeded),
			),
		),
		// 负载均衡，使用 consul 作服务发现
		grpc.WithBalancer(grpc.RoundRobin(grpclb.NewConsulResolver(
			"127.0.0.1:8500", "grpc.health.v1.add",
		))),
		// grpc.WithBalancer(grpc.RoundRobin(grpclb.NewPseudoResolver(
		// 	[]string{
		// 		"127.0.0.1:3000",
		// 		"127.0.0.1:3001",
		// 		"127.0.0.1:3002",
		// 	},
		// ))),
	)
	if err != nil {
		fmt.Printf("dial failed. err: [%v]\n", err)
		return
	}
	defer conn.Close()

	client := addservice.NewAddServiceClient(conn)

	// 限流器
	// 每 800ms 产生 1 个 token，最多缓存 1 个 token，如果缓存满了，新的 token 会被丢弃
	limiter := rate.NewLimiter(rate.Every(time.Duration(800)*time.Millisecond), 1)

	// 熔断器
	hystrix.ConfigureCommand(
		"addservice", // 熔断器名字，可以用服务名称命名，一个名字对应一个熔断器，对应一份熔断策略
		hystrix.CommandConfig{
			Timeout:                100,  // 超时时间 100ms
			MaxConcurrentRequests:  2,    // 最大并发数，超过并发返回错误
			RequestVolumeThreshold: 4,    // 请求数量的阀值，用这些数量的请求来计算阀值
			ErrorPercentThreshold:  25,   // 错误率阀值，达到阀值，启动熔断，25%
			SleepWindow:            1000, // 熔断尝试恢复时间，1000ms
		},
	)

	for i := int64(1); i < 10; i++ {
		for j := int64(1); j < 10; j++ {
			// 限流
			if err := limiter.Wait(context.Background()); err != nil {
				panic(err)
			}

			{
				// 熔断，阻塞方式调用
				var res *addservice.AddResponse
				req := &addservice.AddRequest{A: i, B: j}
				err := hystrix.Do("addservice", func() error {
					// 正常业务逻辑，一般是访问其他资源
					var err error
					// 设置总体超时时间 10 ms 超时
					ctx, cancel := context.WithTimeout(context.Background(), time.Duration(50*time.Millisecond))
					defer cancel()
					res, err = client.Add(
						ctx, req,
						// 这里可以再次设置重试次数，重试时间，重试返回码
						grpc_retry.WithMax(3),
						grpc_retry.WithPerRetryTimeout(time.Duration(5)*time.Millisecond),
						grpc_retry.WithCodes(codes.DeadlineExceeded),
					)
					return err
				}, func(err error) error {
					// 失败处理逻辑，访问其他资源失败时，或者处于熔断开启状态时，会调用这段逻辑
					// 可以简单构造一个response返回，也可以有一定的策略，比如访问备份资源
					// 也可以直接返回 err，这样不用和远端失败的资源通信，防止雪崩
					// 这里因为我们的场景太简单，所以我们可以在本地在作一个加法就可以了
					fmt.Println(err)
					res = &addservice.AddResponse{V: req.A + req.B}
					return nil
				})

				// 事实上这个断言永远为假，因为错误会触发熔断调用 fallback，而 fallback 函数返回 nil
				if err != nil {
					fmt.Printf("client add failed. err: [%v]\n", err)
				}

				fmt.Println(req, res)
			}

			{
				// 熔断，非阻塞方式调用
				var res1 *addservice.AddResponse
				var res2 *addservice.AddResponse
				req := &addservice.AddRequest{A: i, B: j}
				success := make(chan struct{}, 2)

				// 无 fallback 处理
				errc1 := hystrix.Go("addservice", func() error {
					var err error
					ctx, cancel := context.WithTimeout(context.Background(), time.Duration(50*time.Millisecond))
					defer cancel()
					res1, err = client.Add(ctx, req)
					if err == nil {
						success <- struct{}{}
					}
					return err
				}, nil)

				// 有 fallback 处理
				errc2 := hystrix.Go("addservice", func() error {
					var err error
					ctx, cancel := context.WithTimeout(context.Background(), time.Duration(50*time.Millisecond))
					defer cancel()
					res2, err = client.Add(ctx, req)
					if err == nil {
						success <- struct{}{}
					}
					return err
				}, func(err error) error {
					fmt.Println(err)
					res2 = &addservice.AddResponse{V: req.A + req.B}
					success <- struct{}{}
					return nil
				})

				for i := 0; i < 2; i++ {
					select {
					case <-success:
						fmt.Println("success", i)
					case err := <-errc1:
						fmt.Println("err1:", err)
					case err := <-errc2:
						// 这个分支永远不会走到，因为熔断机制里面永远不会返回错误
						fmt.Println("err2:", err)
					}
				}

				fmt.Println(req, res1, res2)
			}
		}
	}
}
