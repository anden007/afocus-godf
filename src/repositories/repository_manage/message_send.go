package repository_manage

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type MessageSendRepository struct{}

func NewMessageSendRepository() *MessageSendRepository {
	instance := new(MessageSendRepository)
	return instance
}

func (m *MessageSendRepository) Add(entity model_manage.MessageSend) (err error) {
	db := interfaces.DI().GetDataBase()

	err = db.GetDB().Create(&entity).Error
	return
}

func (m *MessageSendRepository) Edit(entity model_manage.MessageSend) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Save(&entity).Error
	return
}

func (m *MessageSendRepository) DelByIds(ids []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Delete(model_manage.MessageSend{}, "id in (?)", ids).Error
	return
}

func (m *MessageSendRepository) GetById(id uuid.UUID) (result model_manage.MessageSend, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.MessageSend{}).Where(&model_manage.MessageSend{Id: id}).Preload("Message").First(&result).Error
	return
}

func (m *MessageSendRepository) GetAll() (result []model_manage.MessageSend, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.MessageSend{}).Preload("Message").Find(&result).Error
	return
}

func (m *MessageSendRepository) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.MessageSend{}).Where(model_manage.MessageSend{Id: id}).Updates(fieldValues).Error
	return
}

func (m *MessageSendRepository) GetByCondition(condition map[string]string) (result []model_manage.MessageSend, total int64, err error) {
	db := interfaces.DI().GetDataBase()
	Query := db.GetDB().Model(model_manage.MessageSend{})
	CountQuery := db.GetDB().Model(model_manage.MessageSend{})
	err = lib.NewQueryCondition().GetQuery(model_manage.MessageSend{}.TableName(), condition, Query, CountQuery)
	if err == nil {
		err = CountQuery.Count(&total).Error
		err = Query.Preload("Message").Find(&result).Error
	}
	return
}
