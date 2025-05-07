# 队列使用方法

- 实现接口：`github.com/ilaziness/gokit/queue/rocketmq/Consumer`
- 按需在入口处导入业务队列：`import _ "gintpl/internal/queue"`
- 加载消费组模块：`rocketmq.InitConsumer(web.Config.RocketMq)`
