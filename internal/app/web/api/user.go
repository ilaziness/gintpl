// Package api 定义接口的输入和输出结构
package api

type UserCreateReq struct {
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UserGetReq struct {
	ID int `uri:"id" binding:"required"`
}
type UserGetRes struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Username string `json:"username"`
}
