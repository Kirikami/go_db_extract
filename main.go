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

	ch := make(chan string)
	for _, config := range configs {
		go fetchDatabase(db, ch)
		db := ConnectToNewDatabase(config)
		UserTableDataProvider(db, config)
		SalesTableDataProvider(db, config)
	}
}

func fetchDatabase(db *sqlx.DB, ch chan<- string) {
	start := time.Now()
	UserTableDataProvider(db)
	SalesTableDataProvider(db)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs, %s", secs, db)
}
