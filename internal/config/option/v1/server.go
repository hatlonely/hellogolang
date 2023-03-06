package server

type Server struct {
	host string
	port int
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

func NewServer(opts ...Option) *Server {
	server := &Server{
		host: "localhost",
		port: 3306,
	}
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func Usage() {
	_ = NewServer(WithHost("localhost"), WithPort(3306))
}
