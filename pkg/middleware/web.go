package middleware

import (
	"gintpl/pkg/log"
	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(_ *gin.Context) {
		log.Logger.Info("test middleware")
	}
}
