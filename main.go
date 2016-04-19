package main

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
	_ "time"

	"github.com/kirikami/go_db_extract/config"
	"github.com/kirikami/go_db_extract/database"
	"github.com/kirikami/go_db_extract/services"
)

func main() {
	configfile := flag.String("Configfile", "config.toml", "a string")
	flag.Parse()
	configs := MustNewConfig(configfile)
	ch := make(chan string)
	for _, config := range configs {
		go func() {
			db := MustNewDatabase(config)
			fetchDatabase(db, ch)
		}()
	}
}

func fetchDatabase(db *sqlx.DB, ch chan<- string) {
	start := time.Now()
	UserTableDataProvider(db)
	SalesTableDataProvider(db)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs, %s", secs, db)
}
