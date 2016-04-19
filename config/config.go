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
	Db_name  string
	FilePath string
}

func ReadConfig() Config {
	var configfile = ("../config.toml")
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
