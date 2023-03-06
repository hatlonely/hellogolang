package server

import (
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func NewServer(host string, port int) *Server {
	return &Server{host: host, port: port, timeout: time.Minute, maxConn: 10}
}

func NewServerWithTimeout(host string, port int, timeout time.Duration) *Server {
	return &Server{host: host, port: port, timeout: timeout, maxConn: 10}
}

func NewServerWithTimeoutAndMaxConn(host string, port int, timeout time.Duration, maxConn int) *Server {
	return &Server{host: host, port: port, timeout: timeout, maxConn: maxConn}
}

func Usage() {
	_ = NewServer("localhost", 3306)
	_ = NewServerWithTimeout("localhost", 3306, 2*time.Minute)
	_ = NewServerWithTimeoutAndMaxConn("localhost", 3306, 2*time.Minute, 20)
}
