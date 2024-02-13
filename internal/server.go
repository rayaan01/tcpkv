package internal

import (
	"fmt"
	"net"
)

type Server struct {
	networkType   string
	listenAddress string
	listener      net.Listener
}

func CreateServer(address string) Server {
	return Server{
		networkType:   "tcp",
		listenAddress: address,
	}
}

func (s *Server) StartServer() {
	ln, err := net.Listen(s.networkType, s.listenAddress)
	if err != nil {
		HandleError("Failed to start server", err)
	}
	s.listener = ln
	fmt.Println("Server is listening on:", s.listenAddress)
	s.acceptConnections()
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			HandleError("Failed to accept connection", err)
		}
		go HandleConnection(conn)
	}
}
