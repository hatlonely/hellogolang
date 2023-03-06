package server

import (
	"time"
)

type Options struct {
	Host string
	Port int
	// 1. 新增配置项
	Timeout time.Duration
	MaxConn int
}

type Server struct {
	options *Options
}

// 2. 修改原有构造函数，新增默认值
func NewServerWithOptions(options *Options) *Server {
	if options.Timeout == 0 {
		options.Timeout = time.Minute
	}
	if options.MaxConn == 0 {
		options.MaxConn = 10
	}

	return &Server{
		options: options,
	}
}

// 3. 调用处调整
func Usage() {
	_ = NewServerWithOptions(&Options{
		Host: "localhost",
		Port: 3306,
	})
	_ = NewServerWithOptions(&Options{
		Host:    "localhost",
		Port:    3306,
		Timeout: 2 * time.Minute,
	})
	_ = NewServerWithOptions(&Options{
		Host:    "localhost",
		Port:    3306,
		Timeout: 2 * time.Minute,
		MaxConn: 20,
	})
}
