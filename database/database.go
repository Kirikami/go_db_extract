package database

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kirikami/go_db_extract/config"
)

type User struct {
	UserID int    `db:"user_id"`
	Name   string `db:"name"`
}

type Seller struct {
	OrderID     int     `db:"order_id"`
	UserID      int     `db:"user_id"`
	OrderAmount float64 `db:"order_amount"`
}

func MustNewDatabase(c config.Config) *sqlx.DB {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", c.Username, c.Password, c.Host, c.Port, c.DbName)
	db, err := sqlx.Open("mysql", dbConnection)
	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)
	}

	return db
}
