package main

import (
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	logger := log.New(f, "", log.LstdFlags)
	srv := NewBuilder("localhost", 8888).
		Timeout(time.Minute).
		Logger(logger).
		Build()

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	param serverParam
}

type serverParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewBuilder(host string, port int) *serverParam {
	return &serverParam{host: host, port: port}
}
func (sb *serverParam) Timeout(timeout time.Duration) *serverParam {
	sb.timeout = timeout
	return sb
}
func (sb *serverParam) Logger(logger *log.Logger) *serverParam {
	sb.logger = logger
	return sb
}

func (sb *serverParam) Build() *Server {
	svr := &Server{
		param: *sb}
	return svr
}

func (s *Server) Start() error {
	if s.param.logger != nil {
		s.param.logger.Println("server started")
	}
	// do something
	return nil
}
