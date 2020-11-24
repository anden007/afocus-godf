package repository_manage

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"

	"time"
)

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	instance := new(TaskRepository)
	return instance
}

func (m *TaskRepository) Add(entity model_manage.Task) (err error) {
	db := interfaces.DI().GetDataBase()

	if entity.CreateTime.IsZero() {
		entity.CreateTime = time.Now()
	}

	err = db.GetDB().Create(&entity).Error
	return
}

func (m *TaskRepository) Edit(entity model_manage.Task) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Save(&entity).Error
	return
}

func (m *TaskRepository) DelByIds(ids []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Delete(model_manage.Task{}, "id in (?)", ids).Error
	return
}

func (m *TaskRepository) GetById(id uuid.UUID) (result model_manage.Task, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Task{}).Where(&model_manage.Task{Id: id}).First(&result).Error
	return
}

func (m *TaskRepository) GetAll() (result []model_manage.Task, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Task{}).Find(&result).Error
	return
}

func (m *TaskRepository) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Task{}).Where(model_manage.Task{Id: id}).Updates(fieldValues).Error
	return
}

func (m *TaskRepository) GetByCondition(condition map[string]string) (result []model_manage.Task, total int64, err error) {
	db := interfaces.DI().GetDataBase()
	Query := db.GetDB().Model(model_manage.Task{})
	CountQuery := db.GetDB().Model(model_manage.Task{})
	err = lib.NewQueryCondition().GetQuery(model_manage.Task{}.TableName(), condition, Query, CountQuery)
	if err == nil {
		err = CountQuery.Count(&total).Error
		err = Query.Find(&result).Error
	}
	return
}
