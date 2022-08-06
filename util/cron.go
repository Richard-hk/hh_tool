package util

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func SetCronTask(spec string, cmd func()) {
	c := cron.New()
	_, err := c.AddFunc(spec, cmd)
	if err != nil {
		panic(fmt.Sprintf("setCronTask err %s", err))
	}
	c.Start()
}

func SetCronTaskWithSecond(spec string, cmd func()) {
	c := cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
	_, err := c.AddFunc(spec, cmd)
	if err != nil {
		panic(fmt.Sprintf("setCronTask err %s", err))
	}
	c.Start()
}
