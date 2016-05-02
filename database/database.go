package database

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kirikami/go_db_extract/config"
)

var (
	ErrDbConnect = errors.New("Failed connect to database")
)

func NewDatabase(c config.Config) (*sqlx.DB, error) {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", c.Username, c.Password, c.Host, c.Port, c.DbName)
	db, err := sqlx.Connect("mysql", dbConnection)
	if err != nil {
		return nil, ErrDbConnect
	}

	return db, nil
}

func MustNewDatabase(c config.Config) *sqlx.DB {
	db, err := NewDatabase(c)
	if err != nil {
		log.Fatalf("Connection problem: %s", err)
	}
	return db
}
