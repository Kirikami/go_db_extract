package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	_ "time"

	"github.com/kirikami/go_db_extractor/config"
	"github.com/kirikami/go_db_extractor/database"
)

type Users struct {
	User_ID int    `sql:"AUTO_INCREMENT"`
	Name    string `sql:"varchar(255)"`
}

type Sales struct {
	Order_ID     int     `sql:"AUTO_INCREMENT"`
	User_ID      int     `sql:"type:int(10)"`
	Order_amount float64 `sql:"type:float(50)"`
}

func main() {

	var config Config
	var dbConnection string

	source, err := ioutil.ReadFile("config.yml")

	if err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalf("Cannot parce yaml: %s", err)
	}

	dbConnection = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Db_name)
	db, err := sqlx.Open("mysql", dbConnection)
	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)

	//start := time.Now()
	//	ch := make(chan string)
	//	for _, tablename := range os.Args[1:] {
	dumpTable(db, "users")
	//	}
}

func dumpTable(db *sqlx.DB, tablename string) {

	if tablename == "users" {
		fmt.Println("Chosen users")
		db.Select("*").Find(&users)
		for _, value := range users {
			stringToWrite := []string{strconv.Itoa(value.User_ID), value.Name}
			fmt.Print(stringToWrite)
			err := writer.Write(stringToWrite)
			checkError("Cannot write to file", err)
		}

	} else if tablename == "sales" {
		fmt.Println("Chosen sales")
		db.Select("*").Find(&sales)
		for _, value := range sales {
			stringToWrite := []string{strconv.Itoa(value.Order_ID), strconv.Itoa(value.User_ID), strconv.FormatFloat(value.Order_amount, 'f', 6, 64)}
			err := writer.Write(stringToWrite)
			checkError("Cannot write to file", err)
		}
	}

	defer writer.Flush()
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func UserTableDataProvider(db *sqlx.DB) {
	record := fetchDatabase(db)
	
}

func SalesTableDataProvider() {

}
