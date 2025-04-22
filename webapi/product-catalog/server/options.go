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

func (s *Server) ReadTimeout(t time.Duration) {
	s.httpserver.ReadTimeout = t
}

func (s *Server) WriteTimeout(t time.Duration) {
	s.httpserver.WriteTimeout = t
}

func (s *Server) ShutdownTimeout(t time.Duration) {
	s.shutdownTimeout = t
}
