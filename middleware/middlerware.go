package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Middleware struct {
	Log *zap.Logger
}

func NewMiddleware(log *zap.Logger) Middleware {
	return Middleware{
		Log: log,
	}
}

func (middleware *Middleware) MiddlewareLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Proses request ke handler berikutnya
		c.Next()

		// Menghitung durasi request
		duration := time.Since(start)

		// Logging informasi request
		middleware.Log.Info("http request",
			zap.String("url", c.Request.URL.String()),
			zap.String("method", c.Request.Method),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
		)
	}
}
