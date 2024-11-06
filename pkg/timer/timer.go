package timer

import (
	"gintpl/pkg/log"

	"github.com/robfig/cron/v3"
)

var jobs []Jober
var scheduler *cron.Cron

// running 是否有启动定时器
var running bool

// Jober 任务接口
type Jober interface {
	cron.Job
	GetName() string
	GetCron() string
}

// RegisterJob 注册任务
func RegisterJob(job Jober) {
	jobs = append(jobs, job)
}

// Run 启动定时器
func Run() {
	if len(jobs) == 0 {
		return
	}
	scheduler = cron.New()
	for _, job := range jobs {
		scheduler.AddJob(job.GetCron(), job)
	}
	scheduler.Start()
	running = true
	log.Logger.Infoln("timer started")
}

// Stop 停止定时器
func Stop() {
	if !running {
		return
	}
	log.Logger.Infoln("timer stop")
	ctx := scheduler.Stop()
	<-ctx.Done()
	log.Logger.Infoln("timer stop done")
}
