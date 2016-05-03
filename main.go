package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	_ "github.com/jmoiron/sqlx"
	"github.com/kirikami/go_db_extract/config"
	"github.com/kirikami/go_db_extract/workers"
	"time"
)

func main() {
	start := time.Now()
	configfile := flag.String("config", "config.toml", "Config for connection to database")
	flag.Parse()
	configs := config.MustNewConfig(configfile)
	workers.DbWork(configs)
	secs := time.Since(start).Seconds()
	log.Infof("Program finishing dump in %0.3fs", secs)
}
