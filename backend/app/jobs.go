package app

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/golang-module/carbon/v2"
)

func StartJobs(cfg *config.ServerConfig) {

	dbEngine, err := db.NewEngine(cfg.Db)
	if err != nil {
		panic(err)
	}

	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	s.NewJob(gocron.DurationJob(time.Duration(cfg.JobsConfig.DeviceCheckJob.Duration)*time.Second), gocron.NewTask(func() {
		expired := carbon.Now(cfg.Db.TimeZone).SubSeconds(30).ToDateTimeString()
		dbEngine.Where("is_online = 1 and updated_at <= ?", expired).Cols("is_online").Update(&model.Device{
			IsOnline: false,
		})
	}))

	s.Start()
}
