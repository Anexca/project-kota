package workers

import (
	"log"

	"github.com/robfig/cron/v3"
)

func InitWorkers() *cron.Cron {
	c := cron.New()

	c.AddFunc("*/1 * * * *", func() {
		log.Println("Every Minute")
	})

	return c
}
