package repository_manage

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type DepartmentRepository struct{}

func NewDepartmentRepository() *DepartmentRepository {
	instance := new(DepartmentRepository)
	return instance
}

func (m *DepartmentRepository) GetByParentId(parentId uuid.UUID) (result []model_manage.Department, err error) {
	db := interfaces.DI().GetDataBase()
	if parentId == uuid.Nil {
		err = db.GetDB().Model(&model_manage.Department{}).Where("parent_id is null or parent_id = ?", uuid.Nil).Preload("Parent").Find(&result).Error
	} else {
		err = db.GetDB().Model(&model_manage.Department{}).Where("parent_id = ?", parentId).Preload("Parent").Find(&result).Error
	}
	return
}

func (m *DepartmentRepository) Add(entity model_manage.Department) (err error) {
	db := interfaces.DI().GetDataBase()
	if entity.ParentId != uuid.Nil {
		err = db.GetDB().Model(model_manage.Department{}).Where(model_manage.Department{Id: entity.ParentId}).Updates(model_manage.Department{IsParent: true}).Error
	}
	err = db.GetDB().Create(&entity).Error
	return
}

func (m *DepartmentRepository) Edit(entity model_manage.Department) (err error) {
	db := interfaces.DI().GetDataBase()
	if entity.ParentId != uuid.Nil {
		err = db.GetDB().Model(model_manage.Department{}).Where(model_manage.Department{Id: entity.ParentId}).Updates(model_manage.Department{IsParent: true}).Error
	}
	err = db.GetDB().Save(&entity).Error
	return
}

func (m *DepartmentRepository) DelByIds(ids []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Delete(model_manage.Department{}, "id in (?)", ids).Error
	if err == nil {
		err = db.GetDB().Exec("CALL func_sync_isparent").Error
	}
	return
}
