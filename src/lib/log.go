package lib

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

// 创建日志文件
func GetLogFile() *os.File {
	_, err := os.Stat("./logs")
	if !os.IsExist(err) {
		// 文件夹不存在则创建
		_ = os.Mkdir("./logs", 0666)
	}
	filename := "./logs/main.log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

//"disable" < "fatal" < "error" < "warn" < "info" < "debug"
func parseLogLevel(level string) (result logrus.Level) {
	switch level {
	case "fatal":
		result = logrus.FatalLevel
	case "error":
		result = logrus.ErrorLevel
	case "warn":
		result = logrus.WarnLevel
	case "info":
		result = logrus.InfoLevel
	case "debug":
		result = logrus.DebugLevel
	}
	return
}

type ES_LogCenter struct {
	Log *logrus.Logger
}

func (m *ES_LogCenter) Init() {
	var err error
	logLevel := os.Getenv("log-level")
	logMode := os.Getenv("log-mode")
	appName := os.Getenv("app-name")
	elasticUrl := os.Getenv("elastic-url")
	serverName := os.Getenv("GODF_SERVER_NAME")
	if serverName == "" {
		serverName = "Unknow"
	}
	m.Log = logrus.New()
	m.Log.Formatter = new(logrus.JSONFormatter)
	m.Log.Level = parseLogLevel(logLevel)
	if IS_DEV_MODE {
		m.Log.Formatter = new(logrus.TextFormatter)
		m.Log.Out = os.Stdout
	} else {
		if strings.EqualFold(logMode, "elastic") {
			if client, clientErr := elastic.NewClient(elastic.SetURL(elasticUrl), elastic.SetSniff(false)); clientErr == nil {
				if hook, hookErr := elogrus.NewAsyncElasticHook(client, serverName, m.Log.Level, fmt.Sprintf("log-%s", strings.ToLower(appName))); hookErr == nil {
					m.Log.Hooks.Add(hook)
				} else {
					err = hookErr
				}
			} else {
				err = clientErr
			}
			if err != nil {
				logFile := GetLogFile()
				m.Log.Out = logFile
				m.Log.Error("elastic log init failed! use file log mode.")
			}
		} else {
			logFile := GetLogFile()
			m.Log.Out = logFile
		}
	}
}

var LogCenterInstance *ES_LogCenter
var LogCenterSingle sync.Once

func LogCenter() *logrus.Logger {
	LogCenterSingle.Do(
		func() {
			LogCenterInstance = new(ES_LogCenter)
			LogCenterInstance.Init()
		})
	return LogCenterInstance.Log
}
