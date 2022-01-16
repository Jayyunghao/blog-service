package middleware

import (
	"Practice/go-programming-tour-book/blog-service/pkg/app"
	"Practice/go-programming-tour-book/blog-service/pkg/errcode"
	"Practice/go-programming-tour-book/blog-service/pkg/limiter"

	"github.com/gin-gonic/gin"
)

func RateLimiter (l limiter.LimiterInterface)  gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return 
			}
		}
		c.Next() 
	}
}