package server

import (
	"time"
)

type Options struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
}

type Server struct {
	options *Options
}

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
