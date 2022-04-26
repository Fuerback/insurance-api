package server

import (
	"fmt"
	"net/http"
	"useorigin.com/insurance-api/config/env"
	"useorigin.com/insurance-api/internal/httpadapter/insurancehttpadapter"
)

// Server struct
type Server struct {
	router  Router
	handler insurancehttpadapter.InsuranceHttpHandler
}

// NewServer New Server constructor
func NewServer(handler insurancehttpadapter.InsuranceHttpHandler) *Server {
	return &Server{router: NewMuxRouter(), handler: handler}
}

func (s *Server) Run() {
	s.router.GET("/", func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(resp, "Server up and running...")
	})
	s.router.POST("/evaluation", s.handler.EvaluateUserProfile)

	s.router.Serve(env.GetEnvWithDefaultAsString("PORT", ":8000"))
}
