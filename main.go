package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/kirikami/go_db_extract/config"
	"github.com/kirikami/go_db_extract/database"
	"github.com/kirikami/go_db_extract/services"
	"time"
)

func main() {
	configfile := flag.String("config", "config.toml", "Config for connection to database")
	flag.Parse()
	configs := config.MustNewConfig(configfile)
	ch := make(chan string)
	for _, config := range configs.Database {
		go func() {
			db := database.MustNewDatabase(config)
			archiveDatabase(db, config, ch)
		}()
	}
}

func archiveDatabase(db *sqlx.DB, c config.Config, ch chan<- string) {
	start := time.Now()
	err := services.UserTableDataProvider(db, c)
	if err != nil {
		log.Fatalf("Failed to dump user database: %s", err)
	}
	err = services.SalesTableDataProvider(db, c)
	if err != nil {
		log.Fatalf("Failed to dump sales database: %s", err)
	}
	err = services.ArchiveFile(c.FilePath, ".")
	if err != nil {
		log.Fatalf("Archieving failed: %s", err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs, %s", secs, db)
}
