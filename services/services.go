package services

import (
	"archive/tar"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirikami/go_db_extract/config"
	"github.com/kirikami/go_db_extract/database"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	ErrCantReadFile = errors.New("Cant read file")
)
var (
	ErrCantCreateFile = errors.New("Cannot create file")
)
var (
	ErrCantWriteFile = errors.New("Cannot write to file")
)
var (
	ErrCantFetchData = errors.New("Failed to fetch data")
)
var (
	ErrCantCreateDirectory = errors.New("Cant create directory")
)

type Result struct {
	DbName     string
	FinishTime float64
}

var writeReult []string

func generateCSV(tablename, filepath string, records [][]string) error {
	folder, err := folderExists(filepath)
	if folder != true {
		os.Mkdir(filepath, 0777)
	}
	if err != nil {
		return ErrCantCreateDirectory
	}
	file, err := os.Create(filepath + "/" + tablename + ".csv")
	if err != nil {
		return ErrCantCreateFile
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, stringToWrite := range records {
		err := writer.Write(stringToWrite)
		if err != nil {
			return ErrCantWriteFile
		}
	}
	defer writer.Flush()
	return nil
}

func UserTableDataProvider(db *sqlx.DB, c config.Config) error {
	users := []database.User{}
	err := db.Select(&users, "SELECT * FROM users;")
	if err != nil {
		return ErrCantFetchData
	}
	records := make([][]string, len(users)-1)
	for _, record := range users {
		records = append(records, []string{strconv.Itoa(record.UserID), record.Name})
	}
	records = append(records, []string{fmt.Sprintf("There are %d records in database", len(users))})
	err = generateCSV("users", c.FilePath, records[1:])
	if err != nil {
		return err
	}
	return nil
}

func SalesTableDataProvider(db *sqlx.DB, c config.Config) error {
	sales := []database.Seller{}
	err := db.Select(&sales, "SELECT * FROM sales;")
	if err != nil {
		return ErrCantFetchData
	}
	records := make([][]string, len(sales)-1)
	for _, record := range sales {
		records = append(records, []string{strconv.Itoa(record.OrderID), strconv.Itoa(record.UserID), strconv.FormatFloat(record.OrderAmount, 'f', 6, 64)})
	}
	records = append(records, []string{fmt.Sprintf("There are %d records in database", len(sales))})
	err = generateCSV("sales", c.FilePath, records[1:])
	if err != nil {
		return err
	}
	return nil
}

func ArchiveFile(source, target string) error {
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

func folderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func ArchiveDatabase(db *sqlx.DB, c config.Config, res chan<- Result, errors chan<- error) {
	start := time.Now()
	result := Result{}
	err := UserTableDataProvider(db, c)
	if err != nil {
		errors <- fmt.Errorf("Failed to dump user database: %s", err)
	}
	err = SalesTableDataProvider(db, c)
	if err != nil {
		errors <- fmt.Errorf("Failed to dump sales database: %s", err)
	}
	err = ArchiveFile(c.FilePath, ".")
	if err != nil {
		errors <- fmt.Errorf("Archieving failed: %s", err)
	}
	secs := time.Since(start).Seconds()
	result.DbName = c.DbName
	result.FinishTime = secs
	res <- result
}
