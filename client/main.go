package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"kriyss.ninja/Cow5/client/config"
)

var (
	Error *log.Logger
	Info  *log.Logger
)

func main() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	conf := loadConfiguration("./config.json")
	conn := connect(conf.AddressPort())

	for {
		reader := bufio.NewReader(os.Stdin)
		Info.Print("Message to send : ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		Info.Printf("Message from server: %s", string(message))
	}
}

func loadConfiguration(src string) config.Configuration {
	conf, err := config.Load(src)
	if err != nil {
		Error.Println("Error loading configuration")
		Error.Printf(err.Error())
	}
	loadArgs(conf)
	return *conf
}

func loadArgs(c *config.Configuration) {
	args := os.Args
	for i, arg := range args {
		switch {
		case arg == "--address", arg == "-a":
			c.Server.Address = args[i+1]
		case arg == "--port", arg == "-p":
			c.Server.Port = args[i+1]
		}
	}
}

func connect(address string) net.Conn {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic("Error while connecting : " + err.Error())
	}
	// TODO bug quand je defer la fermeture de connexion.
	// defer conn.Close()
	Info.Printf("Connected to server : %s", conn.RemoteAddr())
	return conn
}
