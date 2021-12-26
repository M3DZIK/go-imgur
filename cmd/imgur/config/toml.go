package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Client ConfigClient
}

type ConfigClient struct {
	ID string
}

func ParseConfig() Config {
	var config Config

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("Get config dir: " + err.Error())
	}

	_, err = toml.DecodeFile(fmt.Sprintf("%s/imgur/config.toml", configDir), config)
	if err != nil {
		panic("Decode config file: " + err.Error())
	}

	return config
}
