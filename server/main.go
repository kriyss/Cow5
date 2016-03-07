package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

var (
	info *log.Logger
)

func main() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	ln, _ := net.Listen("tcp", ":8081")
	info.Printf("Launching server on %s", ln.Addr())

	conn, _ := ln.Accept()
	info.Printf("Connexion accepted from %s", conn.RemoteAddr())

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		info.Printf("Message Received: %s", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage))
	}
}
