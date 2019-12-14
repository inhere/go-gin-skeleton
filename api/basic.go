package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inhere/go-gin-skeleton/app"
)

// JsonData is api response body structure. HttpRes
type JsonData struct {
	Code    int         `json:"code" example:"0" description:"return status code, 0 is successful."`
	Message string      `json:"message" description:"return message"`
	Data    interface{} `json:"data" description:"return data"`
}

// JsonListData only use for swagger docs
type JsonListData struct {
	Code    int      `json:"code" description:"return code, 0 is successful."`
	Message string   `json:"message" description:"return message"`
	Data    []string `json:"data" description:"return data"`
}

// JsonMapData only use for swagger docs
type JsonMapData struct {
	Code    int    `json:"code" description:"return code, 0 is successful."`
	Message string `json:"message" description:"return message"`

	Data map[string]string `json:"data" description:"return data"`
}

// BaseApi controller
type BaseApi struct {
	lang string
}

// getPageAndSize get and format page, size params
func (a *BaseApi) getPageAndSize(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", app.PageSizeStr)

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	return app.FormatPageAndSize(page, size)
}

// DataRes response json data
func (a *BaseApi) DataRes(data interface{}) *JsonData {
	return a.MakeRes(0, nil, data)
}

// MakeRes
// code custom error code
// empty map:
// 	c.DataRes(map[string]string{})
// empty list:
// 	c.DataRes([]int{})
// err  real error message, the message will not output, only write to log file.
func (a *BaseApi) MakeRes(code int, err error, data interface{}) *JsonData {
	if data == nil {
		// data = map[string]string{}
		data = []string{}
	}

	// get output message by error code e.g err-1201
	friendlyMsg := app.Dtr(fmt.Sprintf("err-%d", code))

	// log and print error message
	if err != nil {
		app.Logger.Warn(fmt.Sprintf("detected response error. code:%d message: %s", code, err.Error()))

		// if open debug
		if app.Debug {
			data = map[string]string{"debug_msg": err.Error()}
		}
	}

	return &JsonData{code, friendlyMsg, data}
}
