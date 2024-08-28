package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port               int
	TokenGeneratorFunc func() (string, error)
	PrivateKeyFunc     func() (string, error)
}

func CreateServer(
	port int,
	tokenGeneratorFunc func() (string, error),
	privateKeyFunc func() (string, error),
) *Server {
	if port == 0 {
		port = 8080
	}

	return &Server{
		Port:               port,
		TokenGeneratorFunc: tokenGeneratorFunc,
		PrivateKeyFunc:     privateKeyFunc,
	}

}

func (s *Server) StartServer() error {
	http.Handle("/token", http.HandlerFunc(s.tokenEndpoint))
	http.Handle("/jwk", http.HandlerFunc(s.jwkEndpoint))

	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}
