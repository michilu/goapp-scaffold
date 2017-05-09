package restapi

import (
	"net/http"
)

// NewServer creates a new api example server but does not configure it
func NewServer(_ interface{}) *Server {
	return new(Server)
}

// ConfigureAPI configures the API and handlers.
func (s *Server) ConfigureAPI() {
}

// ConfigureFlags configures the additional flags defined by the handlers. Needs to be called before the parser.Parse
func (s *Server) ConfigureFlags() {
}

// Server for the example API
type Server struct {
	handler http.Handler
}

// GetHandler returns a handler useful for testing
func (s *Server) GetHandler() http.Handler {
	return s.handler
}
