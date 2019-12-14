package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	// boot and init some services(log, cache, eureka)
	"github.com/inhere/go-gin-skeleton/app"
	"github.com/inhere/go-gin-skeleton/model/mongo"
	"github.com/inhere/go-gin-skeleton/model/myrds"
	"github.com/inhere/go-gin-skeleton/model/mysql"
)

func init()  {
	app.Bootstrap("./config")

	// - redis, mongo, mysql connection
	myrds.InitRedis()
	mysql.InitMysql()
	mongo.InitMongo()
	// initEurekaService()
}

func main() {
	listenSignals()

	// init services
	log.Printf("======================== Begin Running(PID: %d) ========================", os.Getpid())

	// default is listen and serve on 0.0.0.0:8080
	app.Run()
}

// listenSignals Graceful start/stop server
func listenSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go handleSignals(sigChan)
}

// handleSignals handle process signal
func handleSignals(c chan os.Signal) {
	log.Print("Notice: System signal monitoring is enabled(watch: SIGINT,SIGTERM,SIGQUIT)\n")

	switch <-c {
	case syscall.SIGINT:
		fmt.Println("\nShutdown by Ctrl+C")
	case syscall.SIGTERM: // by kill
		fmt.Println("\nShutdown quickly")
	case syscall.SIGQUIT:
		fmt.Println("\nShutdown gracefully")
		// do graceful shutdown
	}

	// sync logs
	_ = app.Logger.Sync()
	_ = mysql.Db().Close()
	mongo.Conn().Close()

	// unregister from eureka
	// erkServer.Unregister()

	// 等待一秒
	time.Sleep(1e9 / 2)
	fmt.Println("\nGoodBye...")

	os.Exit(0)
}
