package crontab

import (
	"fmt"
	"github.com/aqie123/superAdminCore/crontab"
)

func Crontab(crontab *crontab.Crontab) {

	//crontab.

	crontab.NewSchedule().EveryMinute().Function(func() {

		fmt.Println("每分钟")

	})

	crontab.NewSchedule().EveryMinuteAt(2).Function(func() {

		fmt.Println("每两分钟")
	})

}
