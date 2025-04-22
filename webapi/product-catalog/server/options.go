package server

import (
	"net"
	"time"
)

type Option func(*Server)

func Addr(port string) Option {
	return func(s *Server) {
		s.httpserver.Addr = net.JoinHostPort("", port)
	}
}

func ReadTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.httpserver.ReadTimeout = t
	}
}

func WriteTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.httpserver.WriteTimeout = t
	}
}

func ShutdownTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = t
	}
}
