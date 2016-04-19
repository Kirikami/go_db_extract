package services

import (
	"archive/tar"
	"bytes"
	"encoding/csv"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/kirikami/go_db_extract/config"
	"github.com/kirikami/go_db_extract/database"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	ErrCantReadFile = errors.New("Cant read file")
)

func generateCSV(tablename, filepath string, records []string) err {
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
	err := writer.Write("There are %d records in databse")
	if err != nil {
		log.Fatalf("Cannot write to file: %s", err)
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

func archiveFile(source, target string) error {
	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.tar", filename))
	tarfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}
