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

// 2. 修改原有构造函数，新增默认值
func NewServer(host string, port int) *Server {
	return &Server{host: host, port: port, timeout: time.Minute, maxConn: 10}
}

// 3. 新增构造函数，支持新的可配置项
func NewServerWithTimeout(host string, port int, timeout time.Duration) *Server {
	return &Server{host: host, port: port, timeout: timeout, maxConn: 10}
}

func NewServerWithTimeoutAndMaxConn(host string, port int, timeout time.Duration, maxConn int) *Server {
	return &Server{host: host, port: port, timeout: timeout, maxConn: maxConn}
}

func Usage() {
	// 4. 调用处调整
	_ = NewServer("localhost", 3306)
	_ = NewServerWithTimeout("localhost", 3306, 2*time.Minute)
	_ = NewServerWithTimeoutAndMaxConn("localhost", 3306, 2*time.Minute, 20)
}
