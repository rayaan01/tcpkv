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
		panic("Failed to start server:" + err.Error())
	}
	s.listener = ln
	fmt.Println("Server is listening on:", s.listenAddress)
	s.acceptConnections()
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic("Failed to accept connection:" + err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		fmt.Println(string(buffer))
	}
}
