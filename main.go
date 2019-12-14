package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	// boot and init some services(log, cache, eureka)
	"github.com/inhere/go-gin-skeleton/app"

	// init redis, mongo, mysql connection
	_ "github.com/inhere/go-gin-skeleton/model/mongo"
	_ "github.com/inhere/go-gin-skeleton/model/mysql"
	_ "github.com/inhere/go-gin-skeleton/model/rds"

	"github.com/inhere/go-gin-skeleton/web"
	"github.com/inhere/go-gin-skeleton/web/middleware"
)

func main() {
	var r *gin.Engine
	env := os.Getenv("APP_ENV")

	if env == app.DEV {
		r = gin.Default()
	} else {
		r = gin.New()
	}

	// global middleware
	r.Use(middleware.RequestLog())

	web.AddRoutes(r)

	log.Printf("======================== Begin Running(PID: %d) ========================", os.Getpid())

	// default is listen and serve on 0.0.0.0:8080
	r.Run(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
}
