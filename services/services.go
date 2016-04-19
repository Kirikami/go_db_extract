package services

import (
	"encoding/csv"
	log "github.com/Sirupsen/logrus"
	"github.com/kirikami/go_db_extract/config"
	"github.com/kirikami/go_db_extract/database"
	"os"
	"strconv"
	"strings"
)

func generateCSV(tablename, filepath string, records []string) {
	file, err := os.Create(filepath + tablename + ".csv")
	if err != nil {
		log.Fatalf("Cannot create file: %s", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, stringToWrite := range records {
		err := writer.Write(strings.Split(stringToWrite))
		if err != nil {
			log.Fatalf("Cannot write to file: %s", err)
		}
	}
	defer writer.Flush()
}

func UserTableDataProvider(db *sqlx.DB, c *Config) {
	user := Users{}
	users := []User{}
	var records []string
	err = db.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Fatalf("Failed to fetch data: %s", err)
	}
	for _, record := range users {
		records := append(records, strconv.Itoa(record.UserID), record.Name)
	}
	generateCSV("users", c.FilePath, records)

}

func SalesTableDataProvider(db *sqlx.DB, c *Config) {
	seller := Seller{}
	sales := []Seller{}
	var records []string
	err := db.Select(&sales, "SELECT * FROM sales")
	if err != nil {
		log.Fatalf("Failed to fetch data: %s", err)
	}
	for _, record := range users {
		records := append(records, strconv.Itoa(value.OrderID), strconv.Itoa(value.UserID), strconv.FormatFloat(value.OrderAmount, 'f', 6, 64))
	}
	generateCSV("sales", c.FilePath, records)

}

func archiveFile(filename string) {

}
