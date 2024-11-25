package middleware

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"

	"gintpl/pkg/log"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
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

		latency := time.Since(start)
		if latency > time.Minute {
			latency = latency.Truncate(time.Second)
		}
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		log.Logger.Infow(path, "latency", latency.String(), "client_ip", clientIP, "method", method, "status_code", statusCode, "query", raw, "error", errorMessage)
	}
}

func Otel(serviceName string) gin.HandlerFunc {
	isSetProvider := false
	tp := otel.GetTracerProvider()
	trf := reflect.TypeOf(tp)
	if trf.Elem().String() != "global.tracerProvider" {
		isSetProvider = true
	}
	return func(c *gin.Context) {
		if !isSetProvider {
			c.Next()
			return
		}

		ctx := otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		// Tracer 使用相同的参数多次调用，返回相同Tracer
		ctx, span := tp.Tracer(serviceName).Start(ctx, c.Request.RequestURI, oteltrace.WithSpanKind(oteltrace.SpanKindServer))
		defer span.End()

		c.Request = c.Request.WithContext(ctx)
		//otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(c.Writer.Header()))
		c.Writer.Header().Set("Trace-ID", span.SpanContext().TraceID().String())

		c.Next()

		status := c.Writer.Status()
		span.SetStatus(codes.Ok, "")
		if status > 0 {
			span.SetAttributes(attribute.Int("http.status_code", status))
		}
		if len(c.Errors) > 0 {
			span.SetAttributes(attribute.String("gin.errors", c.Errors.String()))
		}
	}
}
