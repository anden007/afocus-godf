package interface_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type IMessageSendService interface {
	Add(entity model_manage.MessageSend) (err error)
	DelByIds(ids []uuid.UUID) (err error)
	Edit(entity model_manage.MessageSend) (err error)
	Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error)
	GetAll() (result []model_manage.MessageSend, err error)
	GetById(id uuid.UUID) (result model_manage.MessageSend, err error)
	GetByCondition(condition map[string]string) (result []model_manage.MessageSend, total int64, err error)
	GetMsgCenterChannelId(userId uuid.UUID) (result string, err error)
}
