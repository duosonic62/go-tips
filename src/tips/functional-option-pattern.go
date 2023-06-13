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
	svr := New(
		"localhost",
		8888,
		WithLogger(logger),
		WithTimeout(time.Minute),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

// Option serverの設定を行う関数を用意してNewでOptionとして渡せるようにする
type Option func(*Server)

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}
func WithLogger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

func New(host string, port int, options ...Option) *Server {
	srv := &Server{
		host: host,
		port: port,
	}

	for _, option := range options {
		option(srv)
	}

	return srv
}

func (s *Server) Start() error {
	// do something
	return nil
}
