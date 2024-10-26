package handler

import (
	"gintpl/pkg/base/response"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	response.Success(c, gin.H{"status": "index page"})
}

func Painc(c *gin.Context) {
	panic("1234")
}
