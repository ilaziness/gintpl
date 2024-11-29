package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ilaziness/gintpl/internal/app/web"
	"github.com/ilaziness/gintpl/internal/dao"
	"github.com/ilaziness/gintpl/internal/ent"
	"github.com/ilaziness/gintpl/internal/errcode"
	"github.com/ilaziness/gokit/base/response"
	"github.com/ilaziness/gokit/log"
	"github.com/ilaziness/gokit/otel"
	"github.com/ilaziness/gokit/queue/rocketmq"
	"github.com/ilaziness/gokit/server"
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

func ServiceDis(c *gin.Context) {
	ip, err := server.GetInstance("GinTpl3")
	log.Info(c, "ServiceDis: %v - %v", ip, err)
	response.Success(c, gin.H{"status": "service dis"})
}

func TestEnt(c *gin.Context) {
	err := dao.User.Create(c, &ent.User{
		Age:      18,
		Name:     c.Query("name"),
		Username: c.Query("username"),
	})

	if err != nil {
		response.Error(c, errcode.CodeDBCreateFailed)
		return
	}
	response.Success(c, nil)
}
