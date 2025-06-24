package router

import (
	"fmt"
	"net/http"

	"github.com/WBTechSchool/L0/internal/http/handlers"
	"github.com/WBTechSchool/L0/internal/http/middleware"
	"github.com/WBTechSchool/L0/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(s *service.Service) *gin.Engine {
	r := gin.New()
	h := handlers.NewHandlers(s)

	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middleware.Logger())

	orders := r.Group("/order")
	{
		orders.GET("/:id", h.OrderHandlers.GetOrderByID)
	}

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
