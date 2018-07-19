package server

import (
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	Logger       *zap.Logger
	Name         string
	UpstreamHost string
}

func (s *Server) Serve(addr string) error {
	m := http.NewServeMux()
	m.HandleFunc("/time", s.timeHandler)
	m.HandleFunc("/echo", s.echoHandler)
	m.HandleFunc("/", s.recurseHandler)
	s.Logger.Info("listening",
		zap.String("address", addr),
		zap.String("name", s.Name),
		zap.String("upstream", s.UpstreamHost))
	return http.ListenAndServe(addr, m)
}
