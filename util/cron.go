package util

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

// CheckCron 检查cron表达式是否正确
func CheckCron(cronTab string) (string, error) {
	var (
		specParser cron.Parser
		sched      cron.Schedule
		err        error
	)
	specParser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	if sched, err = specParser.Parse(cronTab); err != nil {
		return "", err
	}
	return fmt.Sprintf("解析正确, 下次运行时间: %s", sched.Next(time.Now())), nil
}
