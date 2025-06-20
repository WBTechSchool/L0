package router

import (
	"fmt"
	"net/http"

	"github.com/WBTechSchool/L0/internal/http/handlers"
	"github.com/WBTechSchool/L0/internal/http/middleware"
	"github.com/WBTechSchool/L0/internal/service"
	"github.com/gin-gonic/gin"
)

func New(s *service.Service) *gin.Engine {
	r := gin.New()
	handlers.NewHandlers(s)

	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": map[string]any{
				"code":    http.StatusNotFound,
				"message": fmt.Sprintf("Resourse not found: %s", ctx.Request.URL.Path),
			},
		})
	})

	return r
}
