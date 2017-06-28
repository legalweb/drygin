package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func LimitRequestSize(size int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength > size {
			AbortBasicError(c, http.StatusRequestEntityTooLarge, "Request size to large.")
		}
	}
}

func LimitRequestSizeKB(size int64) gin.HandlerFunc {
	return LimitRequestSize(size * (1 << 10))
}
