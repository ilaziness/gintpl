package rocketmq

import (
	"context"
	"errors"
	"os"
	"time"

	"gintpl/pkg/config"
	"gintpl/pkg/log"
	"gintpl/pkg/utils"
	"golang.org/x/sync/semaphore"

	rmq "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

var (
	consumerStarted   bool
	consumers         []Consumer
	consumerCtxCancel context.CancelFunc

	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s，影响消费失败的重试间隔，重试间隔 = invisibleDuration - 消费处理时长
	invisibleDuration = time.Second * 30
)

// Consumer 消费者接口
type Consumer interface {
	// GroupName 组名称
	GroupName() string
	// Number 消费组消费者数量，消费者数量大于1时会异步执行，否则时同步执行
	Number() int
	// Subscribe 订阅的主题和tag, key是要订阅的topic，value是要订阅的tag，所有是*
	Subscribe() map[string]string
	// Run 消费者消费消息逻辑，异步执行时处理成功需要主动调用AckFn函数，否则AckFun是nil
	Run(*rmq.MessageView, AckFn) error
}

// AckFn 处理成功的ack通知
type AckFn func()

// RegisterConsumer 注册消费组组
func RegisterConsumer(c Consumer) {
	consumers = append(consumers, c)
}

// InitConsumer 初始化消费组
func InitConsumer(cfg *config.RocketMq) {
	if len(consumers) == 0 {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	consumerCtxCancel = cancel
	_ = os.Setenv("rocketmq.client.logRoot", "log")
	rmq.ResetLogger()
	for _, cg := range consumers {
		utils.SafeGo(func() {
			StartConsumer(ctx, cfg.Endpoint, cfg.AccessKey, cfg.SecretKey, cg)
		})
	}
	consumerStarted = true
}

// StartConsumer 启动消费者
// rocket mq的原则，同一个消费者组下所有消费者实例所订阅的Topic、Tag必须完全一致，否则可能会出现混乱，丢失
// ctx 仅用来控制退出
func StartConsumer(ctx context.Context, endpoint, accessKey, secretKey string, consumer Consumer) {
	sub := map[string]*rmq.FilterExpression{}
	for topic, tag := range consumer.Subscribe() {
		sub[topic] = rmq.NewFilterExpression(tag)
	}
	sc, err := rmq.NewSimpleConsumer(&rmq.Config{
		Endpoint:      endpoint,
		ConsumerGroup: consumer.GroupName(),
		Credentials: &credentials.SessionCredentials{
			AccessKey:    accessKey,
			AccessSecret: secretKey,
		},
	},
		rmq.WithAwaitDuration(awaitDuration),
		// 设置订阅主题和过滤条件
		rmq.WithSubscriptionExpressions(sub),
	)
	if err != nil {
		log.Logger.Errorf("StartConsumer fail: %v", err)
		panic(err)
	}

	err = sc.Start()
	if err != nil {
		log.Logger.Errorf("StartConsumer start fail: %v", err)
		panic(err)
	}

	// 停止consumer
	defer func() {
		if err = sc.GracefulStop(); err != nil {
			log.Logger.Errorf("consumer group [%s] stop error: %v", consumer.GroupName(), err)
		}
	}()

	// 消费者数量
	gNumber := consumer.Number()
	cNumberPool := semaphore.NewWeighted(int64(gNumber))
	pullMessage := func() {
		newCtx := context.TODO()
		var rpcErr *rmq.ErrRpcStatus
		mvs, err := sc.Receive(newCtx, maxMessageNum, invisibleDuration)
		if err != nil && errors.As(err, &rpcErr) && rpcErr.Code == 40401 {
			time.Sleep(time.Second * 2)
			return
		}
		if err != nil {
			log.Logger.Errorf("consumer receive msg error: %v", err)
			time.Sleep(time.Second * 5)
			return
		}
		for _, mv := range mvs {
			if gNumber == 1 {
				err = consumer.Run(mv, nil)
				if err != nil {
					log.Logger.Errorf("message consume fail: <%s>, %v", mv.GetMessageId(), err)
					continue
				}
				if err = sc.Ack(newCtx, mv); err != nil {
					log.Logger.Errorf("ack fail: <%s>, %v", mv.GetMessageId(), err)
				}
				continue
			}
			// 消费组消费者数量大于1，异步执行
			if err = cNumberPool.Acquire(newCtx, 1); err != nil {
				log.Logger.Errorf("acquire fail: <%s>, %v", mv.GetMessageId(), err)
				continue
			}
			utils.SafeGo(func() {
				_ = consumer.Run(mv, func() {
					defer cNumberPool.Release(1)
					if err = sc.Ack(newCtx, mv); err != nil {
						log.Logger.Errorf("ack fail: <%s>, %v", mv.GetMessageId(), err)
					}
				})
			})
		}
		time.Sleep(time.Second * 3)
	}

	log.Logger.Infof("consumer group started: %s", consumer.GroupName())

	for {
		select {
		case <-ctx.Done():
			return
		default:
			pullMessage()
		}
	}
}

// ConsumerStop 停止所有消费组
func ConsumerStop() {
	if !consumerStarted {
		return
	}
	consumerCtxCancel()
	time.Sleep(time.Second * 3)
	log.Logger.Info("queue consumer stop done")
}
