package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gookit/ini"
	"github.com/inhere/go-gin-skeleton/app/utils"
	"github.com/inhere/go-gin-skeleton/model"

	"github.com/inhere/go-gin-skeleton/app/cache"
)

// Cfg application config
var Cfg *ini.Ini

func init() {
	initAppEnv()

	loadAppConfig()

	log.Printf(
		"======================== Bootstrap (Env: %s, Debug: %v) ========================",
		Env, Debug,
	)

	initAppInfo()

	initLogger()

	initLanguage()

	initCache()

	// initEurekaService()

	listenSignals()
}

func initAppEnv() {
	Hostname, _ = os.Hostname()

	if env := os.Getenv("APP_ENV"); env != "" {
		Env = env
	}

	if port := os.Getenv("APP_PORT"); port != "" {
		HttpPort, _ = strconv.Atoi(port)
	}

	// in dev, test
	if IsEnv(DEV) || IsEnv(TEST) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// loadAppConfig
func loadAppConfig() {
	var err error

	envFile := "conf/app-" + Env + ".ini"

	fmt.Printf("- work dir: %s\n", WorkDir)
	fmt.Printf("- load config: conf/app.ini, %s\n", envFile)

	Cfg, err = ini.LoadFiles("conf/app.ini", envFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Cfg.WriteTo(os.Stdout)

	// setting some info
	Name = Cfg.MustString("name")
	Debug = Cfg.MustBool("debug")
}

func initAppInfo() {
	// ensure http port
	if HttpPort == 0 {
		HttpPort = Cfg.MustInt("httpPort")
	}

	// git repo info
	GitInfo = model.GitInfoData{}
	infoFile := "static/app.json"

	if utils.FileExists(infoFile) {
		utils.ReadJsonFile(infoFile, &GitInfo)
	}
}

// init redis connection pool
func initCache() {
	conf, _ := Cfg.StringMap("cache")

	// 从配置文件获取redis的ip以及db
	prefix := conf["prefix"]
	server := conf["server"]
	password := conf["auth"]
	redisDb, _ := strconv.Atoi(conf["db"])

	fmt.Printf("cache - server=%s db=%d auth=%s\n", server, redisDb, password)

	// 建立连接池
	// closePool()
	cache.Init(NewRedisPool(server, password, redisDb), prefix, Logger, Debug)
}
