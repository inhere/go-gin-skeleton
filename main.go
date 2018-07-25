package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inhere/go-gin-skeleton/route"
	"log"
	"os"
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

	route.AddRoutes(r)

	log.Printf("======================== Begin Running(PID: %d) ========================", os.Getpid())

	// default is listen and serve on 0.0.0.0:8080
	r.Run(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
}
