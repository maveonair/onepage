package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type server struct {
	router       *mux.Router
	pageFilePath string
}

func NewServer(pageFilePath string) (Server, error) {
	server := &server{
		router:       mux.NewRouter(),
		pageFilePath: pageFilePath,
	}

	if err := server.routes(); err != nil {
		return nil, err
	}

	return server, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
