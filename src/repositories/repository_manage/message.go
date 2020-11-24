package repository_manage

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"

	"time"
)

type MessageRepository struct{}

func NewMessageRepository() *MessageRepository {
	instance := new(MessageRepository)
	return instance
}

func (m *MessageRepository) Add(entity model_manage.Message) (err error) {
	db := interfaces.DI().GetDataBase()

	if entity.CreateTime.IsZero() {
		entity.CreateTime = time.Now()
	}

	err = db.GetDB().Create(&entity).Error
	return
}

func (m *MessageRepository) Edit(entity model_manage.Message) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Save(&entity).Error
	return
}

func (m *MessageRepository) DelByIds(ids []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Delete(model_manage.Message{}, "id in (?)", ids).Error
	return
}

func (m *MessageRepository) GetById(id uuid.UUID) (result model_manage.Message, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Message{}).Where(&model_manage.Message{Id: id}).First(&result).Error
	return
}

func (m *MessageRepository) GetAll() (result []model_manage.Message, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Message{}).Find(&result).Error
	return
}

func (m *MessageRepository) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Message{}).Where(model_manage.Message{Id: id}).Updates(fieldValues).Error
	return
}

func (m *MessageRepository) GetByCondition(condition map[string]string) (result []model_manage.Message, total int64, err error) {
	db := interfaces.DI().GetDataBase()
	Query := db.GetDB().Model(model_manage.Message{})
	CountQuery := db.GetDB().Model(model_manage.Message{})
	err = lib.NewQueryCondition().GetQuery(model_manage.Message{}.TableName(), condition, Query, CountQuery)
	if err == nil {
		err = CountQuery.Count(&total).Error
		err = Query.Find(&result).Error
	}
	return
}
