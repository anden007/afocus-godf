package model_manage

import (
	"time"

	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type Task struct {
	model.DBModel
	Id          uuid.UUID `gorm:"size:36;primaryKey" json:"id"`
	FuncName    string    `gorm:"size:100;comment:'任务方法名称'" json:"funcName"`
	Success     bool      `gorm:"default:0;comment:'是否执行成功'" json:"success"`
	Result      string    `gorm:"type:text" json:"result"`
	CreatorId   uuid.UUID `gorm:"size:36;comment:'创建者Id'" json:"creatorId"`
	CreateTime  time.Time `gorm:"size:100" json:"createTime"`
	LastRunTime time.Time `gorm:"size:100" json:"lastRunTime"`
}

func (Task) TableName() string {
	return "t_task"
}
