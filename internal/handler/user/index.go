package user

import (
	"fish/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Index(c *middleware.Context) (any, error) {
	return gin.H{"status": "user page"}, nil
}
