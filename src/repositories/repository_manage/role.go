package repository_manage

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type RoleRepository struct{}

func NewRoleRepository() *RoleRepository {
	instance := new(RoleRepository)
	return instance
}

func (m *RoleRepository) Add(entity model_manage.Role) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Create(&entity).Error
	return
}

func (m *RoleRepository) Edit(entity model_manage.Role) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Save(&entity).Error
	return
}

func (m *RoleRepository) DelByIds(ids []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Delete(model_manage.Role{}, "id in (?)", ids).Error
	return
}

func (m *RoleRepository) EditRoleDep(roleId uuid.UUID, dataType int, depIds []uuid.UUID) error {
	return nil
}

func (m *RoleRepository) GetAll() (result []model_manage.Role, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Role{}).Find(&result).Error
	return
}

func (m *RoleRepository) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Role{}).Where(model_manage.Role{Id: id}).Updates(fieldValues).Error
	return
}

func (m *RoleRepository) EditRolePerm(roleId uuid.UUID, permIds []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	var newPerm []model_manage.Permission
	err = db.GetDB().Model(&model_manage.Permission{}).Where("id in (?)", permIds).Find(&newPerm).Error
	if err == nil {
		var role model_manage.Role
		err = db.GetDB().First(&role, "id = ?", roleId).Error
		err = db.GetDB().Model(&role).Where("id = ?", roleId).Association("Permissions").Replace(&newPerm)
	}
	return
}

func (m *RoleRepository) GetByCondition(condition map[string]string) (result []model_manage.Role, total int64, err error) {
	db := interfaces.DI().GetDataBase()
	Query := db.GetDB().Model(model_manage.Role{})
	CountQuery := db.GetDB().Model(model_manage.Role{})
	err = lib.NewQueryCondition().GetQuery(model_manage.Role{}.TableName(), condition, Query, CountQuery)
	if err == nil {
		err = CountQuery.Count(&total).Error
		err = Query.Preload("Permissions").Preload("Departments").Find(&result).Error
	}
	return
}
