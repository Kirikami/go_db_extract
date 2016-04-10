package main

import (
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	_ "time"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Db_name  string
}

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
	checkError("Cannot read config", err)

	err = yaml.Unmarshal(source, &config)
	checkError("Cannot parce yaml", err)

	dbConnection = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Db_name)
	db, err := gorm.Open("mysql", dbConnection)
	checkError("failed to connect database", err)

	db.DB()
	db.DB().Ping()
	//start := time.Now()
	//	ch := make(chan string)
	//	for _, tablename := range os.Args[1:] {
	dumpTable(db, "users")
	//	}
}

func dumpTable(db *gorm.DB, tablename string) {
	users := []Users{}
	sales := []Sales{}

	file, err := os.Create(tablename + ".csv")
	checkError("Cannot create file", err)

	defer file.Close()

	writer := csv.NewWriter(file)
	fmt.Println(tablename)
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

func archiveFile(filename string) {

}
