package middleware

import (
	"github.com/gin-gonic/gin"
)

func AbortBasicError(g *gin.Context, code int, err string) {
	g.AbortWithStatus(code)
	IssueBasicError(g, code, err)
}

func IssueBasicError(g *gin.Context, code int, err string) {
	switch g.ContentType() {
	case "application/json":
		g.JSON(code, struct {
			Message string `json:"message"`
		}{
			Message: err,
		})
		break
	default:
		g.String(code, "<h1>Error</h1><p>Message: %s</p>", err)
	}
}
