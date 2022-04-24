package server

import (
	"fmt"
	"net/http"
)

const (
	port = ":8000"
)

// Server struct
type Server struct {
	router Router
}

// NewServer New Server constructor
func NewServer() *Server {
	return &Server{router: NewMuxRouter()}
}

func (s *Server) Run() {
	s.router.GET("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Server up and running...")
	})
	//s.router.POST("/evaluation", service)

	s.router.Serve(port)
}
