package config

import (
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

type tomlConfig struct {
	Database []Config
}

func MustNewConfig(configfile string) tomlConfig {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var tomlconf tomlConfig
	_, err = toml.DecodeFile(configfile, &tomlconf)
	if err != nil {
		log.Fatal(err)
	}
	return tomlconf
}
