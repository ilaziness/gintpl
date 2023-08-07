package handler

import (
	"fish/internal/errcode"
	"fish/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *middleware.Context) (any, error) {
	return gin.H{"status": "ok"}, errcode.CodeNotFound
}
