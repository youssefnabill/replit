
package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()

		log.Printf("%s %s - %d - %v", method, path, status, latency)
	}
}
