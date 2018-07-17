package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	User UserConfig `json:"user"`
}

type UserConfig struct {
	Name  string        `json:"name"`
	Age   int           `json:"age"`
	Slave []SlaveConfig `json:"slave"`
}

type SlaveConfig struct {
	Option bool     `json:"option"`
	Array  []string `json:"array"`
}

func main() {
	var config Config
	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(b))
}
