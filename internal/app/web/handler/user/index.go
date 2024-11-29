package user

import (
	"github.com/ilaziness/gokit/base/response"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Success(c, gin.H{"status": "user page"})
}
