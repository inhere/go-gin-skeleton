package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inhere/go-gin-skeleton/app"
	"github.com/inhere/go-gin-skeleton/app/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

func RequestLog() gin.HandlerFunc {
	skip := map[string]int{
		"/":       1,
		"/health": 1,
		"/status": 1,
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		reqId := utils.GenMd5(fmt.Sprintf("%d", start.Nanosecond()))

		// add reqID to context
		c.Set("reqId", reqId)

		// c.MustBindWith()
		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; ok {
			return
		}

		// log post/put data
		postData := ""
		if c.Request.Method != "GET" {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			postData = string(buf)
		}

		app.Logger.Info(
			"complete request",
			zap.String("req_id", reqId),
			zap.Namespace("context"),
			zap.String("req_date", start.Format("2006-01-02 15:04:05")),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.URL.String()),
			zap.String("client_ip", c.ClientIP()),
			zap.Int("http_status", c.Writer.Status()),
			zap.String("elapsed_time", utils.CalcElapsedTime(start)),
			zap.String("post_data", postData),
		)
	}
}
