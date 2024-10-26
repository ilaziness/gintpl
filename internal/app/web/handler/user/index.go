package user

import (
	"gintpl/pkg/base/response"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Success(c, gin.H{"status": "user page"})
}
