package wsserver

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"time"
)

func (s *Server) handlerWsMessages(ws *websocket.Conn) {
	fmt.Println("New conn to ws with", ws.RemoteAddr())

	s.conns[ws] = true // will need mutex

	s.readLoop(ws)
}

func (s *Server) handleWsDataFlow(ws *websocket.Conn) {
	fmt.Println("New conn to broadcast with", ws.RemoteAddr())
	for {
		payload := fmt.Sprintf("subscription data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := ws.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err", err, "on conn", ws.RemoteAddr())
			continue
		}
		msg := buff[:n]
		fmt.Println(string(msg))
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println(err)
			}
		}(ws)
	}
}
