package internal

import (
	"net"
	"strings"
)

var store = map[string]string{}

func HandleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		args := strings.Fields(string(buffer))
		cmd := strings.ToLower(args[0])
		switch cmd {
		case "set":
			responseBytes := handleSet(args[1], args[2])
			conn.Write(responseBytes)
		case "get":
			responseBytes := handleGet(args[1])
			conn.Write(responseBytes)
		}
	}
}

func handleSet(key string, value string) []byte {
	store[key] = value
	return []byte("OK\n")
}

func handleGet(key string) []byte {
	val, ok := store[key]
	if ok {
		return []byte(val + "\n")
	} else {
		return []byte("(nil)\n")
	}
}
