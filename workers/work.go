package workers

import (
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
	databases := len(configs.Database)
	for i := 0; i < databases; i++ {
		go func() {
			for _, c := range configs.Database {
				dumpDatabase := func() {
					db := database.MustNewDatabase(c)
					services.ArchiveDatabase(db, c)
				}
				pending <- NewWork(dumpDatabase)
			}
		}()
	}

	for i := 0; i < databases; i++ {
		go Worker(pending, done)
	}

	for i := 0; i < databases; i++ {
		<-done
	}
}
