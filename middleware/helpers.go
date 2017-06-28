package middleware

import (
	"net/http"
)

func extractHeader(headers http.Header, name string, defaultValue []string) []string {
	if val, ok := headers[name]; ok {
		return val
	}

	return defaultValue
}
