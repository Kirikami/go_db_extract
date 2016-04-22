package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"os"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
	FilePath string
}

type TomlConfig struct {
	Database []Config
}

var (
	ErrCantFindConfig = errors.New("Config file is missing")
)

func NewConfig(configfile string) (*TomlConfig, error) {
	_, err := os.Stat(configfile)
	if err != nil {
		return nil, ErrCantFindConfig
	}

	var tomlconf *TomlConfig
	_, err = toml.DecodeFile(configfile, &tomlconf)
	if err != nil {
		return nil, err
	}
	return tomlconf, nil
}

func MustNewConfig(configfile *string) *TomlConfig {
	config, err := NewConfig(*configfile)
	if err != nil {
		log.Fatalf("Cant parse config file: %s", err)
	}
	return config
}
