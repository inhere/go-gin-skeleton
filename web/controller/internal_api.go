package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/inhere/go-gin-skeleton/app"
)

// InternalApi
type InternalApi struct {
	BaseApi
}

// @Tags InternalApi
// @Summary Get app config
// @Param	key		query 	string	 false		"config key string"
// @Success 201 {string} json data
// @Failure 403 body is empty
// @router /config [get]
func (a *InternalApi) Config(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		key = app.Config.DefSection()
	}

	val, _ := app.Config.StringMap(key)

	c.JSON(200, val)
}
