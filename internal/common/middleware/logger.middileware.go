package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		logMessage := fmt.Sprintf("[%s] %s %s %d %s",
			method, path, clientIP, statusCode, latency)

		if statusCode >= 500 {
			fmt.Printf("ERROR: %s\n", logMessage)
		} else if statusCode >= 400 {
			fmt.Printf("WARN: %s\n", logMessage)
		} else {
			fmt.Printf("INFO: %s\n", logMessage)
		}
	}
}
