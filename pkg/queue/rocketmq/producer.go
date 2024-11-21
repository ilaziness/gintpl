package rocketmq

import (
	"context"
	"errors"
	"os"

	"gintpl/pkg/config"
	"gintpl/pkg/log"

	rmq "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

var producer rmq.Producer
var producerStarted bool

var ErrReceipt = errors.New("send message receipt empty")

// InitProducer 初始化rocket mq生产者
func InitProducer(cfg *config.RocketMq) {
	var err error
	options := []rmq.ProducerOption{rmq.WithTopics(cfg.ProducerTopic)}
	if cfg.Transaction {
		options = append(options, rmq.WithTransactionChecker(&rmq.TransactionChecker{
			Check: func(msg *rmq.MessageView) rmq.TransactionResolution {
				log.Logger.Infof("check transaction message: %v", msg)
				return rmq.COMMIT
			},
		}))
	}
	os.Setenv("rocketmq.client.logRoot", "log")
	rmq.ResetLogger()
	producer, err = rmq.NewProducer(&rmq.Config{
		Endpoint: cfg.Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    cfg.AccessKey,
			AccessSecret: cfg.SecretKey,
		},
	}, options...)
	if err != nil {
		log.Logger.Errorf("initRMQProducer fail: %v", err)
		panic(err)
	}
	err = producer.Start()
	if err != nil {
		log.Logger.Errorf("initRMQProducer start fail: %v", err)
		panic(err)
	}
	producerStarted = true
	log.Logger.Info("queue producer create done")
}

// ProducerStop 停止生产者
func ProducerStop() {
	if !producerStarted {
		return
	}
	producer.GracefulStop()
	log.Logger.Info("queue producer stop done")
}

type MessageOpts struct {
	Tag       string
	Keys      []string
	Group     string
	DelayTime int // 延迟秒数
}

// Send 发送消息
// rocket mq 客户端内部有重试逻辑，建议调用者要处理发送错误避免发送失败消息丢失，失败时可以把消息写入到本地db
func Send(ctx context.Context, topic string, body []byte, opts ...MessageOpts) error {
	msg := &rmq.Message{
		Topic: topic,
		Body:  body,
	}
	if len(opts) > 0 {
		if opts[0].Tag != "" {
			msg.SetTag(opts[0].Tag)
		}
		if len(opts[0].Keys) > 0 {
			msg.SetKeys(opts[0].Keys...)
		}
		if opts[0].Group != "" {
			msg.SetMessageGroup(opts[0].Group)
		}
	}
	receipt, err := producer.Send(ctx, msg)
	if err != nil {
		return err
	}
	if len(receipt) == 0 || receipt[0].MessageID == "" {
		return ErrReceipt
	}
	return nil
}
