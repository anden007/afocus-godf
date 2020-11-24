package model_manage

import (
	"time"

	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type Message struct {
	model.DBModel
	Id             uuid.UUID   `gorm:"size:36;primaryKey" json:"id"`
	Title          string      `gorm:"size:100" json:"title"`
	Content        string      `gorm:"size:1000" json:"content"`
	MessageType    int         `gorm:"comment:'0:通知 1:私信 2:消息'" json:"messageType"`
	Range          int         `gorm:"comment:'0:全体 1:指定用户'" json:"range"`
	UserIds        []uuid.UUID `gorm:"-" json:"userIds"`
	SenderId       uuid.UUID   `gorm:"size:36" json:"senderId"`
	SenderNickName string      `gorm:"size:100" json:"senderNickName"`
	SendTime       time.Time   `json:"sendTime"`
	CreateTime     time.Time   `json:"createTime"`
}

func (Message) TableName() string {
	return "t_message"
}
