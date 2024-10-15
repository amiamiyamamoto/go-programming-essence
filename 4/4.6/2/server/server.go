package server

import (
	"log"
	"time"
)

type Server struct {
	param ServerParam
}

type ServerParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewBuilder(host string, port int) *ServerParam {
	return &ServerParam{host: host, port: port}
}

func (sb *ServerParam) Timeout(timeout time.Duration) *ServerParam {
	sb.timeout = timeout
	return sb
}
func (sb *ServerParam) Logger(logger *log.Logger) *ServerParam {
	sb.logger = logger
	return sb
}

func (sb *ServerParam) Build() *Server {
	svr := &Server{
		param: *sb,
	}
	return svr
}

func (s *Server) Start() error {
	if s.param.logger != nil {
		s.param.logger.Println("server started")
	}
	//do something
	return nil
}
