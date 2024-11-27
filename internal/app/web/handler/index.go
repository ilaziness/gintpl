package handler

import (
	"gintpl/internal/app/web"
	"gintpl/pkg/base/response"
	"gintpl/pkg/log"
	"gintpl/pkg/otel"
	"gintpl/pkg/queue/rocketmq"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	log.Warn(c, "%+v", web.Config.App)
	log.Info(c, "%+v", web.Config.App)
	log.Debug(c, "%+v", web.Config.App)
	log.Error(c, "%+v", web.Config.App)
	response.Success(c, gin.H{"status": "index page"})
}

func Painc(_ *gin.Context) {
	panic("1234")
}

func SendMq(c *gin.Context) {
	_, span := otel.Tracer.Start(c.Request.Context(), "sendMq")
	defer span.End()
	log.Logger.Info("send mq ", rocketmq.Send(c, "test1", []byte("测试数据")))

	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据2")))
	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据3")))
	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据4")))
	response.Success(c, gin.H{"status": "send mq"})
}

func Trace(c *gin.Context) {
	ctx, span := otel.Tracer.Start(c, "index")
	defer span.End()
	_, span2 := otel.Tracer.Start(ctx, "index2")
	defer span2.End()

	log.Warn(ctx, "hello")

	response.Success(c, gin.H{"text": "Trace demo"})
}
