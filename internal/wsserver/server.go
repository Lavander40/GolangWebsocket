package wsserver

import (
	"golang.org/x/net/websocket"
	"net/http"
)

type Server struct {
	config *Config
	conns  map[*websocket.Conn]bool
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
		conns:  make(map[*websocket.Conn]bool),
	}
}

func (s *Server) Start() error {
	http.Handle("/messages", websocket.Handler(s.handlerWsMessages))
	http.Handle("/dataflow", websocket.Handler(s.handleWsDataFlow))
	if err := http.ListenAndServe(s.config.BindAddr, nil); err != nil {
		return err
	}
	return nil
}
