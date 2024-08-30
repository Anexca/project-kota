package workers

import (
	"log"

	"github.com/robfig/cron/v3"
)

type Worker struct {
	cronHandler *cron.Cron
}

func InitWorkers() *cron.Cron {
	c := cron.New()

	worker := Worker{
		cronHandler: c,
	}

	worker.RegisterWorkers()
	return c
}

func (w *Worker) RegisterWorkers() {
	w.cronHandler.AddFunc("*/1 * * * *", func() {
		log.Println("Every Minute")
	})
}
