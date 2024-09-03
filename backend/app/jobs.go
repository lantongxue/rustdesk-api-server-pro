package app

import (
	"fmt"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func Jobs(cfg *config.ServerConfig) {

	dbEngine, err := db.NewEngine(cfg.Db)
	if err != nil {
		panic(err)
	}

	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	s.NewJob(gocron.DurationJob(time.Duration(cfg.JobsConfig.DeviceCheckJob.Duration)*time.Second), gocron.NewTask(func() {
		t, _ := time.ParseDuration("-30s")
		expired := time.Now().Add(t).Format("2006-01-02 15:04:05")
		fmt.Println(expired)
		dbEngine.Where("is_online = 1 and updated_at <= ?", expired).Cols("is_online").Update(&model.Device{
			IsOnline: false,
		})
	}))

	s.Start()
}
