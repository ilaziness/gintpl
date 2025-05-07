package queue

import (
	"log"

	"github.com/ilaziness/gokit/queue/rocketmq"

	rmq "github.com/apache/rocketmq-clients/golang/v5"
)

type Test1 struct {
	groupName string
}

func init() {
	rocketmq.RegisterConsumer(NewTest1())
}

func NewTest1() *Test1 {
	return &Test1{
		groupName: "test1",
	}
}

func (t1 Test1) GroupName() string {
	return t1.groupName
}

func (t1 Test1) Number() int {
	return 1
}

func (t1 Test1) Subscribe() map[string]string {
	return map[string]string{
		"test1": "*",
	}
}

func (t1 Test1) Run(mv *rmq.MessageView, ack rocketmq.AckFn) error {
	log.Println(t1.GroupName(), mv.GetMessageId(), string(mv.GetBody()))
	if ack != nil {
		ack()
		return nil
	}
	return nil
}
