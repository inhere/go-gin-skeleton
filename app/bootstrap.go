package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/dotnev"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/i18n"
	"github.com/inhere/go-gin-skeleton/web"
	"github.com/inhere/go-gin-skeleton/web/middleware"

	"github.com/inhere/go-gin-skeleton/app/cache"
	// - redis, mongo, mysql services
	"github.com/inhere/go-gin-skeleton/model"
)

var (
	I18n   *i18n.I18n
	Config *config.Config
	server *gin.Engine
)

func Bootstrap(confDir string) {
	initAppEnv()

	loadAppConfig(confDir)

	color.Info.Printf(
		"======================== Bootstrap (EnvName: %s, Debug: %v) ========================",
		EnvName, Debug,
	)

	initAppInfo()

	initLogger()

	initLanguage()

	initCache()
}

func Run() {
	initServer()

	err := server.Run(fmt.Sprintf("0.0.0.0:%d", HttpPort))
	if err != nil {
		color.Error.Println(err)
	}
}

func initAppEnv() {
	err := dotnev.LoadExists(".", ".env")
	if err != nil {
		color.Error.Println(err.Error())
		return
	}

	Hostname, _ = os.Hostname()
	if env := os.Getenv("APP_ENV"); env != "" {
		EnvName = env
	}

	if port := os.Getenv("APP_PORT"); port != "" {
		HttpPort, _ = strconv.Atoi(port)
	}

	// in dev, test
	if IsEnv(EnvDev) || IsEnv(EnvTest) {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// loadAppConfig
func loadAppConfig(confDir string) {
	baseFile := confDir + "app" + configSuffix
	envFile := confDir + "/app-" + EnvName + configSuffix

	fmt.Printf("- work dir: %s\n", WorkDir)
	fmt.Printf("- load config: conf/app.ini, %s\n", envFile)

	err := config.LoadFiles(baseFile, envFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	Config = config.Default()
	// Config.WriteTo(os.Stdout)

	// setting some info
	Name = Config.String("name")
	Debug = Config.Bool("debug")
}

func initAppInfo() {
	// ensure http port
	if HttpPort == 0 {
		HttpPort = Config.Int("httpPort")
	}

	// git repo info
	GitInfo = model.AppInfo{}
	infoFile := "static/app.json"

	if fsutil.IsFile(infoFile) {
		err := jsonutil.ReadFile(infoFile, &GitInfo)
		if err != nil {
			color.Error.Println(err.Error())
		}
	}
}

// init redis connection pool
func initCache() {
	conf := Config.StringMap("cache")

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

func initServer() {
	server = gin.New()

	if IsEnv(EnvDev) {
		server.Use(gin.Logger(), gin.Recovery())
	}

	// global middleware
	server.Use(middleware.RequestLog())

	web.AddRoutes(server)
}
