package workers

import (
	"time"
)

type Work struct {
	Function  func()
	Completed bool
}

func NewWork(function func()) *Work {
	return &Work(function, false)
}

func Worker(in chan *Work, out chan *Work) {
	work := <-in
	work.Function()
	work.Completed = true
	out <- work
}

func RunWork(function func(), jobs, conc int) {
	pending := make(chan *Work)
	done := make(chan *Work)

	go func() {
		for i := 0; i < jobs; i++ {
			pending <- NewWork(function)
		}
	}()
	for i := 0; i < conc; i++ {
		go Worker(pending, done)
	}

	for i := 0; i < jobs; i++ {
		<-done
	}
}
