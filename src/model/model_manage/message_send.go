package model_manage

import (
	"time"

	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type MessageSend struct {
	model.DBModel
	Id               uuid.UUID `gorm:"size:36;primaryKey" json:"id"`
	MessageId        uuid.UUID `gorm:"size:36" json:"messageId"`
	Status           int       `gorm:"comment:'0:未读 1:已读 2:回收站'" json:"status"`
	Message          *Message  `gorm:"foreignKey:MessageId" json:"message"`
	ReceiverId       uuid.UUID `gorm:"size:36" json:"receiverId"`
	ReceiverNickName string    `gorm:"size:100" json:"receiverNickName"`
	ReadTime         time.Time `json:"readTime"`
	CreateTime       time.Time `json:"createTime"`
}

func (MessageSend) TableName() string {
	return "t_messagesend"
}
