package web

import (
	"github.com/gin-gonic/gin"
	"github.com/inhere/go-gin-skeleton/web/controller"
)

// AddRoutes
func AddRoutes(r *gin.Engine) {
	r.GET("/", controller.Home)

	r.LoadHTMLFiles("res/views/swagger.tpl")
	r.GET("/api-docs", controller.SwagDoc)

	// status
	r.GET("/health", controller.AppHealth)
	r.GET("/status", controller.AppStatus)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/health", controller.AppHealth)

		internal := new(controller.InternalApi)
		v1.GET("/config", internal.Config)
	}

	// static assets
	r.Static("/static", "./static")

	// not found routes
	r.NoRoute(func(c *gin.Context) {
		c.JSON(
			404,
			controller.JsonMapData{0, "page not found", map[string]string{}},
		)
	})
}
