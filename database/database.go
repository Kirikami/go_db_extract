package database

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kirikami/go_db_extract/config"
)

type User struct {
	UserID int    `sql:"AUTO_INCREMENT"`
	Name   string `sql:"varchar(255)"`
}

type Saller struct {
	OrderID     int     `sql:"AUTO_INCREMENT"`
	UserID      int     `sql:"type:int(10)"`
	OrderAmount float64 `sql:"type:float(50)"`
}

type database struct {
	db *sqlx.DB
}

func MustNewDatabase(c Config) *sqlx.DB {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", c.Username, c.Password, c.Host, c.Port, c.DbName)
	db, err := sqlx.Open("mysql", dbConnection)
	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)
	}

	return db
}
