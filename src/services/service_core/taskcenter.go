package service_core

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anden007/afocus-godf/src/lib"

	ytask "github.com/gojuukaze/YTask/v2"
	"github.com/gojuukaze/YTask/v2/server"
	"github.com/spf13/cast"
)

type TaskCenter struct {
	Server *server.Server
}

func NewTaskCenter() *TaskCenter {
	_ = lib.LoadEnv()
	instance := new(TaskCenter)
	loadTime := time.Now()
	enable := os.Getenv("taskcenter-enable")
	if strings.ToLower(enable) == "true" {
		if lib.IS_DEV_MODE {
			fmt.Println("> Service: TaskCenter loaded.", time.Since(loadTime))
		}
	} else {
		fmt.Printf("> Service: TaskCenter is disabled. if you need enable it,please set 'taskcenter-enable' to 'true' in .env file.\n")
	}
	return instance
}

func (m *TaskCenter) newServer(clientPoolSize int) (result *server.Server) {
	rdsServer := os.Getenv("taskcenter-rds-server")
	rdsPort := os.Getenv("taskcenter-rds-port")
	rdsPassword := os.Getenv("taskcenter-rds-password")
	db := cast.ToInt(os.Getenv("taskcenter-rds-db"))
	broker := ytask.Broker.NewRedisBroker(rdsServer, rdsPort, rdsPassword, db, clientPoolSize)
	backend := ytask.Backend.NewRedisBackend(rdsServer, rdsPort, rdsPassword, db, clientPoolSize)
	tmpServer := ytask.Server.NewServer(
		ytask.Config.Broker(&broker),
		ytask.Config.Backend(&backend),
		ytask.Config.Debug(lib.IS_DEV_MODE),
		ytask.Config.StatusExpires(60*5),
		ytask.Config.ResultExpires(60*5),
	)
	result = &tmpServer
	return
}

func (m *TaskCenter) GetServer() (result server.Server) {
	if m.Server == nil {
		m.Server = m.newServer(0)
	}
	result = *m.Server
	return
}

func (m *TaskCenter) GetClient(clientPoolSize int) (result server.Client) {
	if m.Server == nil {
		m.Server = m.newServer(clientPoolSize)
	}
	result = m.Server.GetClient()
	return
}
