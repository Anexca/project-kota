package workers

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func InitWorkers() *cron.Cron {
	c := cron.New()

	c.AddFunc("*/1 * * * *", func() { fmt.Println("Every hour on the half hour") })
	return c
}
