package workers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/kirikami/go_db_extract/config"
	"github.com/kirikami/go_db_extract/database"
	"github.com/kirikami/go_db_extract/services"
)

type Work struct {
	Function  func()
	Completed bool
}

func NewWork(function func()) *Work {

	return &Work{function, false}

}

func Worker(in chan *Work, out chan *Work) {

	for {
		work := <-in
		work.Function()
		work.Completed = true
		out <- work
	}

}

func DbWork(configs *config.TomlConfig) {
	pending := make(chan *Work)
	done := make(chan *Work)
	errors := make(chan error)
	result := make(chan services.Result, 1)

	databases := len(configs.Database)

	for _, config := range configs.Database {
		c := config
		db := database.MustNewDatabase(c)
		dumpDatabase := func() {
			services.ArchiveDatabase(db, c, result, errors)
		}

		go func() {
			pending <- NewWork(dumpDatabase)
		}()

	}

	for i := 0; i < databases; i++ {

		go Worker(pending, done)

	}

	for i := 0; i < databases; i++ {
		<-done

		select {
		case res := <-result:
			log.Infof("Database %s dump sucessful in %.3fs", res.DbName, res.FinishTime)
		case err := <-errors:
			log.Fatal(err.Error())
		}
	}

}
