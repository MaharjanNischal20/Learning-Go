package http

import (
	"net/http"
	"time"

	"example.com/auth/internal/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(authHandler *auth.Handler, jwtService *auth.JWTService) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})
	api := router.Group("/api/v1")
	apiRoutes := api.Group("/auth")
	{
		apiRoutes.POST("/register", authHandler.Register)
		apiRoutes.POST("/login", authHandler.Login)
	}
	protected := api.Group("")
	protected.Use(jwtService.Middleware())
	{
		protected.GET("/auth/me", authHandler.Me)
	}
	return router
}
