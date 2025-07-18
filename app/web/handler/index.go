package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ilaziness/gintpl/app/web"
	"github.com/ilaziness/gintpl/db/dao"
	"github.com/ilaziness/gintpl/db/ent"
	"github.com/ilaziness/gintpl/errcode"
	"github.com/ilaziness/gokit/base/reqres"
	"github.com/ilaziness/gokit/log"
	"github.com/ilaziness/gokit/otel"
	"github.com/ilaziness/gokit/queue/rocketmq"
	"github.com/ilaziness/gokit/server"
	"github.com/ilaziness/gokit/storage/sql"
)

// Index 首页
func Index(c *gin.Context) {
	reqres.Success(c, gin.H{"status": "ok"})
}

func Painc(_ *gin.Context) {
	panic("1234")
}

func LogTest(c *gin.Context) {
	log.Warn(c, "%+v", web.Config.App)
	log.Info(c, "%+v", web.Config.App)
	log.Debug(c, "%+v", web.Config.App)
	log.Error(c, "%+v", web.Config.App)

	reqres.Success(c, "log test")
}

func SendMq(c *gin.Context) {
	_, span := otel.Tracer.Start(c.Request.Context(), "sendMq")
	defer span.End()
	log.Logger.Info("send mq ", rocketmq.Send(c, "test1", []byte("测试数据")))

	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据2")))
	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据3")))
	log.Logger.Info("send mq ", rocketmq.Send(c, "test2", []byte("测试数据4")))
	reqres.Success(c, gin.H{"status": "send mq"})
}

func Trace(c *gin.Context) {
	ctx, span := otel.Tracer.Start(c, "index")
	defer span.End()
	_, span2 := otel.Tracer.Start(ctx, "index2")
	defer span2.End()

	log.Warn(ctx, "hello")

	reqres.Success(c, gin.H{"text": "Trace demo"})
}

func ServiceDis(c *gin.Context) {
	ip, err := server.GetInstance("GinTpl3")
	log.Info(c, "ServiceDis: %v - %v", ip, err)
	reqres.Success(c, gin.H{"status": "service dis"})
}

type User struct {
	ID        int    `db:"id"`
	Age       int    `db:"age"`
	Name      string `db:"name"`
	Username  string `db:"username"`
	CreatedAt string `db:"created_at"`
}

func TestEnt(c *gin.Context) {
	err := dao.User.Create(c, &ent.User{
		Age:      18,
		Name:     c.Query("name"),
		Username: c.Query("username"),
	})

	if err != nil {
		log.Error(c, "%v", err)
		reqres.Error(c, errcode.CodeDBCreateFailed)
		return
	}

	//u := User{}
	//err = mysql.SqlxDB().Get(&u, "SELECT * FROM users LIMIT 1")
	//log.Info(c, "User: %v - %v", u, err)

	reqres.Success(c, nil)
}

func TestGorm(c *gin.Context) {
	u := User{}
	err := sql.GormDB().WithContext(c).Table("users").First(&u).Error
	log.Info(c, "User: %v - %v", u, err)

	reqres.Success(c, nil)
}
