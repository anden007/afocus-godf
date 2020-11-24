package repository_manage

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type PermissionRepository struct{}

func NewPermissionRepository() *PermissionRepository {
	instance := new(PermissionRepository)
	return instance
}

func (m *PermissionRepository) GetAll() (result []model_manage.Permission, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Permission{}).Order("sort_order").Find(&result).Error
	return
}

func (m *PermissionRepository) Add(entity model_manage.Permission) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Create(&entity).Error
	return
}

func (m *PermissionRepository) Edit(entity model_manage.Permission) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Save(&entity).Error
	return
}

func (m *PermissionRepository) DelByIds(ids []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Delete(model_manage.Permission{}, "id in (?)", ids).Error
	return
}

func (m *PermissionRepository) GetUserPermissions(userId uuid.UUID) (permissions []model_manage.Permission, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Table("t_user u").Select("DISTINCT p.*").Joins("LEFT JOIN t_user_roles ur ON u.id = ur.user_id").Joins("LEFT JOIN t_role_permissions rp ON ur.role_id = rp.role_id").Joins("LEFT JOIN t_permission p ON p.id = rp.permission_id").Where("u.id = ? AND p.status = 0", userId).Order(" p.sort_order asc").Scan(&permissions).Error
	return
}
