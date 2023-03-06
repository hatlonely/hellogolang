package server

type Options struct {
	Host string
	Port int
}

type Server struct {
	options *Options
}

func NewServerWithOptions(options *Options) *Server {
	return &Server{
		options: options,
	}
}

func Usage() {
	_ = NewServerWithOptions(&Options{
		Host: "localhost",
		Port: 3306,
	})
}
