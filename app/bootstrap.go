package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/dotnev"
	"github.com/gookit/config/v2/toml"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/i18n"
	"github.com/inhere/go-gin-skeleton/app/cache"
	// - redis, mongo, mysql services
	"github.com/inhere/go-gin-skeleton/model"
)

var (
	I18n   *i18n.I18n
	Config *config.Config
)

func Bootstrap(confDir string) {
	initAppEnv()

	loadAppConfig(confDir)

	color.Info.Printf(
		"======================== Bootstrap (EnvName: %s, Debug: %v) ========================\n",
		EnvName, Debug,
	)

	initAppInfo()

	initLogger()

	initLanguage()

	initCache()
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
	fmt.Printf("- work dir: %s\n", WorkDir)
	fmt.Printf("- config dir: %s\n", confDir)
	// fmt.Printf("- load config: conf/app.ini, %s\n", envFile)

	files, err := getConfigFiles(confDir)
	if err != nil {
		color.Error.Println(err.Error())
		os.Exit(1)
	}

	// config instance
	Config = config.Default()
	Config.AddDriver(toml.Driver)

	err = Config.LoadFiles(files...)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Config.WriteTo(os.Stdout)
	// setting some info
	Name = Config.String("name")
	Debug = Config.Bool("debug")
}

// 获取某个文件夹下的配置文件列表
func getConfigFiles(confDir string) ([]string, error) {
	var files = make([]string, 0)

	fileInfoList, err := ioutil.ReadDir(confDir)
	if err != nil {
		return files, err
	}

	pathSep := string(os.PathSeparator)
	// app.toml is must exists
	baseFile := confDir + pathSep + "app" + configSuffix
	files = append(files, baseFile)

	// _dev.toml
	suffix := "-" + EnvName + configSuffix
	for _, f := range fileInfoList {
		// app_dev.toml
		if !f.IsDir() && strings.HasSuffix(f.Name(), suffix) {
			files = append(files, confDir+pathSep+f.Name())
		}
	}

	return files, err
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
