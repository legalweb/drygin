package responses

import (
	"net/http"
	"strings"

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

// IssueValidationError will return validation errors.
func IssueValidationError(g *gin.Context, errs []error) {
	sErr := []string{}

	for _, err := range errs {
		sErr = append(sErr, err.Error())
	}

	switch g.ContentType() {
	case "application/json":
		g.JSON(http.StatusBadRequest, struct {
			Messages []string `json:"messages"`
		}{
			Messages: sErr,
		})
		break
	default:
		g.String(
			http.StatusBadRequest,
			"<h1>Bad Request</h1><p>Message: %s</p>",
			strings.Join(sErr, ", "),
		)
	}
}
