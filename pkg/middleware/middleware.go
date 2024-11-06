package middleware

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"gintpl/pkg/log"

	"github.com/gin-gonic/gin"
)

// RecoveryHandle 记录panic日志
func RecoveryHandle(c *gin.Context, err any) {
	message := fmt.Sprintf("%s", err)

	// 记录错误日志
	// log.Logger.Errorw("Panic recovered",
	// 	zap.Time("time", time.Now()),
	// 	zap.Any("error", err),
	// 	zap.String("path", c.Request.URL.Path),
	// 	zap.String("stack", trace(message)),
	// )
	log.Logger.Errorw("panic recovered",
		"error", err,
		"path", c.Request.URL.Path,
		"stack", trace(message),
	)
	c.AbortWithStatus(http.StatusInternalServerError)
}

// trace 获取panic堆栈信息
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(4, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	var str strings.Builder
	str.WriteString(message + "\nTranceback:")
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		str.WriteString(fmt.Sprintf("\n\t%s:%d", frame.File, frame.Line))
	}

	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// LogReq 记录请求日志
func LogReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		latency := time.Now().Sub(start)
		if latency > time.Minute {
			latency = latency.Truncate(time.Second)
		}
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		log.Logger.Infow(path, "latency", fmt.Sprintf("%v", latency), "client_ip", clientIP, "method", method, "status_code", statusCode, "query", raw, "error", errorMessage)
	}
}
