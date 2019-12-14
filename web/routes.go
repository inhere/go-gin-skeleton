package web

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes
func AddRoutes(r *gin.Engine) {
	r.GET("/", Home)

	r.LoadHTMLFiles("res/views/swagger.tpl")
	r.GET("/api-docs", SwagDoc)

	// status
	r.GET("/health", AppHealth)
	r.GET("/status", AppStatus)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/health", AppHealth)

		internal := new(InternalApi)
		v1.GET("/config", internal.Config)
	}

	// static assets
	r.Static("/static", "./static")

	// not found routes
	r.NoRoute(func(c *gin.Context) {
		c.JSON(
			404,
			JsonMapData{0, "page not found", map[string]string{}},
		)
	})
}
