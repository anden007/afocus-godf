package interface_core

import (
	"github.com/gojuukaze/YTask/v2/server"
)

type ITaskCenter interface {
	GetServer() (result server.Server)
	GetClient(clientPoolSize int) (result server.Client)
}
