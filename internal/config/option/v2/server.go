package server

import (
	"time"
)

type Server struct {
	host string
	port int
	// 1. 新增配置项
	timeout time.Duration
	maxConn int
}

type Option func(*Server)

func WithHost(host string) Option {
	return func(server *Server) {
		server.host = host
	}
}

func WithPort(port int) Option {
	return func(server *Server) {
		server.port = port
	}
}

// 2. 新增选项
func WithTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.timeout = timeout
	}
}

func WithMaxConn(maxConn int) Option {
	return func(server *Server) {
		server.maxConn = maxConn
	}
}

// 2. 修改原有构造函数，新增默认值
func NewServer(opts ...Option) *Server {
	server := &Server{
		host:    "localhost",
		port:    3306,
		timeout: time.Minute,
		maxConn: 10,
	}
	for _, opt := range opts {
		opt(server)
	}

	return server
}

// 4. 调用处调整
func Usage() {
	_ = NewServer(
		WithHost("localhost"),
		WithPort(3306),
	)
	_ = NewServer(
		WithHost("localhost"),
		WithPort(3306),
		WithTimeout(time.Minute),
	)
	_ = NewServer(
		WithHost("localhost"),
		WithPort(3306),
		WithTimeout(time.Minute),
		WithMaxConn(20),
	)
}
