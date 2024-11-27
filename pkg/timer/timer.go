package timer

import (
	"gintpl/pkg/hook"
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
	var err error
	scheduler = cron.New()
	for _, job := range jobs {
		if _, err = scheduler.AddJob(job.GetCron(), job); err != nil {
			log.Logger.Warnf("timer add job error, name: %s, err: %v", job.GetName(), err)
		}
	}
	scheduler.Start()
	running = true
	hook.Exit.Register(Stop)
	log.Logger.Infoln("timer started")
}

// Stop 停止定时器
func Stop() {
	if !running {
		return
	}
	ctx := scheduler.Stop()
	<-ctx.Done()
	log.Logger.Infoln("timer stopped")
}
