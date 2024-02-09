package main

import (
	"fmt"
	"net"
)

type Server struct {
	networkType   string
	listenAddress string
	listener      net.Listener
	conn          net.Conn
}

func main() {
	server := createServer("localhost:8080")
	server.startServer()
}

func createServer(address string) Server {
	return Server{
		networkType:   "tcp",
		listenAddress: address,
	}
}

func (s *Server) startServer() {
	ln, err := net.Listen(s.networkType, s.listenAddress)
	if err != nil {
		fmt.Print("Failed to start server: ", err)
	}
	s.listener = ln
	s.acceptConnections()
}

func (s *Server) acceptConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection: ", err)
		}
		s.conn = conn
	}
}

func (s *Server) handleConnection() {
	for {
		buffer := make([]byte, 1024)
		s.conn.Read(buffer)
	}
}
