package queue

import (
	"log"

	"gintpl/pkg/queue/rocketmq"

	rmq "github.com/apache/rocketmq-clients/golang/v5"
)

type Test2 struct {
	groupName string
}

func init() {
	rocketmq.RegisterConsumer(NewTest2())
}

func NewTest2() *Test2 {
	return &Test2{
		groupName: "test2",
	}
}

func (t1 Test2) GroupName() string {
	return t1.groupName
}

func (t1 Test2) Number() int {
	return 3
}

func (t1 Test2) Subscribe() map[string]string {
	return map[string]string{
		"test2": "*",
	}
}

func (t1 Test2) Run(mv *rmq.MessageView, ack rocketmq.AckFn) error {
	log.Println(t1.groupName, mv.GetMessageId(), string(mv.GetBody()))
	ack()
	return nil
}
