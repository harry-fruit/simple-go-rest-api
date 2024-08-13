package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() error {
	fmt.Println("Server is starting on ", s.addr)
	s.startRouters()
	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) startRouters() {
	http.HandleFunc("/", s.handleRoot)
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
