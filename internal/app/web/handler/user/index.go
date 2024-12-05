package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ilaziness/gintpl/internal/app/web/api"
	"github.com/ilaziness/gintpl/internal/app/web/service"
	"github.com/ilaziness/gintpl/internal/errcode"
	"github.com/ilaziness/gokit/base/reqp"
)

func Index(c *gin.Context) {
	reqp.Success(c, gin.H{"status": "user page"})
}

func Create(c *gin.Context) {
	// 用法1
	req := api.UserCreateReq{}
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	reqp.Error(c, errcode.CodeReqErr.SetMessage(err.Error()))
	//	return
	//}
	//if err := service.User.Create(c, &req); err != nil {
	//	reqp.Error(c, err)
	//	return
	//}
	//reqp.Success(c, gin.H{"status": "create user"})

	// 用法2，快捷调用
	reqp.CallServiceWithoutRes(c, req, service.User.Create)
}

func Get(c *gin.Context) {
	req := api.UserGetReq{}
	if err := c.ShouldBindUri(&req); err != nil {
		reqp.Error(c, errcode.CodeReqErr.SetMessage(err.Error()))
		return
	}
	rsp, err := service.User.Get(c, &req)
	if err != nil {
		reqp.Error(c, err)
		return
	}
	reqp.Success(c, rsp)

	// 其他调用sevice的快捷方法
	//reqp.CallService(c, req, service.User.Get)
	//reqp.CallServiceWithoutReqAndRes(c, service.User.List)
	//reqp.CallServiceWithoutReq[[]*api.UserGetRes](c, service.User.List2)
}
