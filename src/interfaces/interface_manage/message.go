package interface_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/google/uuid"
)

type IMessageService interface {
	Add(entity model_manage.Message) (err error)
	DelByIds(ids []uuid.UUID ) (err error)
	Edit(entity model_manage.Message) (err error)
	Updates(id uuid.UUID , fieldValues map[string]interface{}) (err error)
	GetAll() (result []model_manage.Message, err error)
	GetById(id uuid.UUID) (result model_manage.Message, err error)
	GetByCondition(condition map[string]string) (result []model_manage.Message, total int64, err error)
}
