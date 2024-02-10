package main

import "tcpkv/internal"

func main() {
	server := internal.CreateServer("localhost:8080")
	server.StartServer()
}
