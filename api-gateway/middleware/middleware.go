package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(2, 10) 

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}
