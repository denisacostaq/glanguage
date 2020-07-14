package src

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	port uint16
}

func NewServer(port uint16) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	r := mux.NewRouter()
	addr := fmt.Sprintf(":%d", s.port)
	log.WithField("addr", addr).Infoln("Starting server...")
	if err := http.ListenAndServe(addr, r); err != nil {
		return err
	}
	return nil
}
