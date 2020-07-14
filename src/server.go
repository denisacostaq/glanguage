package src

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	port uint16
}

func NewServer(port uint16) *Server {
	return &Server{port: port}
}

func (s *Server) Start() {
	r := mux.NewRouter()
	addr := fmt.Sprintf(":%d", s.port)
	log.Printf("Starting server in %s ...", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalln(err)
	}
}
