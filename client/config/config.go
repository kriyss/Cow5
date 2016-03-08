package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Configuration struct {
	Server struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	}
}

func Load(src string) (*Configuration, error) {
	file, err := os.Open(src)
	if err != nil {
		return &Configuration{}, errors.New("Can't load file at : %s" + src)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Configuration
	err = decoder.Decode(&cfg)
	if err != nil {
		return &Configuration{}, errors.New("Can't decode config file")
	}
	return &cfg, nil
}

func (c *Configuration) AddressPort() string {
	return c.Server.Address + ":" + c.Server.Port
}
