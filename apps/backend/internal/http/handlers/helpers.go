package handlers

import "github.com/gin-gonic/gin"

func JSONErrorWrapper(statusCode int, message string) (int, map[string]any) {
	return statusCode, gin.H{
		"error": map[string]any{
			"code":    statusCode,
			"message": message,
		},
	}
}
