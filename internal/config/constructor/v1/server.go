package server

type Server struct {
	host string
	port int
}

func NewServer(host string, port int) *Server {
	return &Server{host: host, port: port}
}

func Usage() {
	_ = NewServer("localhost", 3306)
}
