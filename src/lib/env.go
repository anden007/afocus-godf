package lib

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

var (
	ENV_LOADED  bool   = false
	IS_DEV_MODE bool   = false
	USE_PPROF   bool   = false
	APP_NAME    string = ""
	APP_HOST    string = ""
	APP_PORT    int    = 0
)

func LoadEnv() (err error) {
	if !ENV_LOADED {
		ENV_LOADED = true
		flag.StringVar(&APP_HOST, "h", "0.0.0.0", "主机IP")
		flag.IntVar(&APP_PORT, "p", 0, "运行端口")
		flag.BoolVar(&USE_PPROF, "d", false, "启用pprof性能测试")
		flag.Parse()

		devMode := os.Getenv("DevMode")
		if devMode == "true" {
			fmt.Println("App Running in Development Mode.")
			IS_DEV_MODE = true
			err = godotenv.Load("development.env")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("App Running in Production Mode.")
			IS_DEV_MODE = false
			err = godotenv.Load("production.env")
			if err != nil {
				log.Fatal(err)
			}
		}
		if APP_PORT == 0 {
			cfgPort := os.Getenv("server-port")
			if cfgPort == "" {
				panic("未设置程序运行端口，无法启动！")
			} else {
				APP_PORT = cast.ToInt(cfgPort)
			}
		}
		APP_NAME = os.Getenv("app-name")
		if USE_PPROF {
			_, err := os.Stat("./debug")
			if !os.IsExist(err) {
				// 文件夹不存在则创建
				_ = os.Mkdir("./debug", 0666)
			}
		}
	}
	return
}
