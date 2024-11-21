# queue用法

## 生产者

1. 应用入口载入组件：

```go
package main

import (
    "gintpl/pkg/queue/rocketmq"
)

rocketmq.InitRMQProducer(cfg)
```

2. 发送消息

```go
package main

import (
    "gintpl/pkg/queue/rocketmq"
)

rocketmq.Send(ctx, "topic", "msg")
```

## 消费者

1. 实现`Consumer`接口
2. 使用`init`注册消费者组
3. 启动消费者：`rocketmq.InitConsumer(web.Config.RocketMq)`

> RocketMQ部署文档：https://rocketmq.io/course/deploy/rocketmq_learning-gvr7dx_awbbpb_ogr2blaw8vy3tv14/?spm=5176.29160081.0.0.50a9608eF04Ez7
