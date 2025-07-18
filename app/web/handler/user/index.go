package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ilaziness/gintpl/app/web/api"
	"github.com/ilaziness/gintpl/app/web/service"
	"github.com/ilaziness/gintpl/errcode"
	"github.com/ilaziness/gokit/base/reqres"
)

func Index(c *gin.Context) {
	reqres.Success(c, gin.H{"status": "user page"})
}

func Create(c *gin.Context) {
	// 用法1
	req := api.UserCreateReq{}
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	reqres.Error(c, errcode.CodeReqErr.SetMessage(err.Error()))
	//	return
	//}
	//if err := service.User.Create(c, &req); err != nil {
	//	reqres.Error(c, err)
	//	return
	//}
	//reqres.Success(c, gin.H{"status": "create user"})

	// 用法2，快捷调用
	reqres.CallServiceNoRes(c, req, service.User.Create)
}

func Get(c *gin.Context) {
	req := api.UserGetReq{}
	if err := c.ShouldBindUri(&req); err != nil {
		errDesc := err.Error()
		reqres.Error(c, errcode.CodeReqErr.SetMessage(errDesc))
		return
	}
	rsp, err := service.User.Get(c, &req)
	if err != nil {
		reqres.Error(c, err)
		return
	}
	reqres.Success(c, rsp)

	// 其他调用sevice的快捷方法
	//reqres.CallService(c, req, service.User.Get)
	//reqres.CallServiceWithoutReqAndRes(c, service.User.List)
	//reqres.CallServiceWithoutReq[[]*api.UserGetRes](c, service.User.List2)
}
