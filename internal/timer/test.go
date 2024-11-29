package timer

import (
	"github.com/ilaziness/gokit/timer"

	"log"
)

func init() {
	timer.RegisterJob(&TestTimer{})
}

type TestTimer struct {
}

func (t *TestTimer) GetName() string {
	return "TestTimer"
}

func (t *TestTimer) GetCron() string {
	return "*/1 * * * *"
}

func (t *TestTimer) Run() {
	log.Println("TestTimer is run")
}
