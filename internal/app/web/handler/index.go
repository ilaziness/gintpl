package handler

import (
	"gintpl/pkg/base/response"
	"gintpl/pkg/log"
	"gintpl/pkg/queue/rocketmq"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	response.Success(c, gin.H{"status": "index page"})
}

func Painc(_ *gin.Context) {
	panic("1234")
}

func SendMq(c *gin.Context) {
	log.Logger.Info("send mq ", rocketmq.Send(c, "test1", []byte("测试数据")))

	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据2")))
	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据3")))
	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据4")))
	response.Success(c, gin.H{"status": "send mq"})
}
