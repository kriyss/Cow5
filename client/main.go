package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	// TODO comment on log de facon en go ?
	logInf *log.Logger
	logErr *log.Logger
)

func main() {
	logInf = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logErr = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	conf := loadConfiguration("config.json")
	conn := connect(conf.AddressPort())

	for {
		reader := bufio.NewReader(os.Stdin)
		logInf.Print("Message to send : ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		logInf.Printf("Message from server: %s", string(message))
	}
}

func loadConfiguration(src string) Configuration {
	conf, err := Load(src)
	if err != nil {
		logErr.Printf("Error loading configuration")
		logErr.Printf(err.Error())
	}
	loadArgs(&conf)
	return conf
}

func loadArgs(c *Configuration) {
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
	logInf.Printf("Connected to server : %s", conn.RemoteAddr())
	return conn
}

// TODO je veux bouger ca dans config.go
type Configuration struct {
	Server Server `json:"server"`
}

type Server struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

func Load(src string) (Configuration, error) {
	file, err := os.Open(src)
	if err != nil {
		return Configuration{}, errors.New("Can't load file at : %s" + src)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Configuration
	err = decoder.Decode(&cfg)
	if err != nil {
		return Configuration{}, errors.New("Can't decode config file")
	}
	return cfg, nil
}

func (c *Configuration) AddressPort() string {
	return c.Server.Address + ":" + c.Server.Port
}
