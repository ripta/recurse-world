package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func (s *Server) echoHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info(r.Method + " " + r.URL.Path)
	io.WriteString(w, "echo says: "+r.URL.RawQuery+"\n")
}

func (s *Server) recurseHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info(r.Method + " " + r.URL.Path)
	stamp := time.Now().Format("2006-01-02T15:04:05.000000Z07:00") + " " + s.Name + " " + r.URL.Path + "\n"

	if r.URL.Path == "/" {
		io.WriteString(w, stamp)
		return
	}

	u := &Upstreamer{
		Host:          s.UpstreamHost,
		OriginalPath:  r.URL.Path,
		OriginalQuery: r.URL.RawQuery,
	}
	body, err := u.Do(r.Context())
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Invalid request: %v", err))
		return
	}
	io.WriteString(w, stamp+body)
}

func (s *Server) timeHandler(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info(r.Method + " " + r.URL.Path)
	io.WriteString(w, "The time is now "+time.Now().Format(time.RFC3339Nano)+"\n")
}
