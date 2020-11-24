package model_manage

import (
	"time"

	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type WxUser struct {
	model.DBModel
	Id         uuid.UUID `gorm:"size:36;primaryKey" json:"id"`
	OpenId     string    `gorm:"size:50"`
	NickName   string    `gorm:"size:50"`
	EMail      string    `gorm:"size:50"`
	RealName   string    `gorm:"size:50"`
	CreateTime time.Time
	PostTime   time.Time
}

func (WxUser) TableName() string {
	return "t_wxuser"
}
