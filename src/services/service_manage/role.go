package service_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/repositories/repository_manage"

	"github.com/google/uuid"
)

type RoleService struct {
	repo *repository_manage.RoleRepository
}

func NewRoleService(repo *repository_manage.RoleRepository) *RoleService {
	instance := &RoleService{
		repo: repo,
	}
	return instance
}
func (m *RoleService) GetAll() (result []model_manage.Role, err error) {
	result, err = m.repo.GetAll()
	return
}

func (m *RoleService) Add(entity model_manage.Role) (err error) {
	err = m.repo.Add(entity)
	return
}

func (m *RoleService) Edit(entity model_manage.Role) (err error) {
	err = m.repo.Edit(entity)
	return
}

func (m *RoleService) DelByIds(ids []uuid.UUID) (err error) {
	err = m.repo.DelByIds(ids)
	return
}

func (m *RoleService) EditRoleDep(roleId uuid.UUID, dataType int, depIds []uuid.UUID) (err error) {
	err = m.repo.EditRoleDep(roleId, dataType, depIds)
	return
}

func (m *RoleService) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	err = m.repo.Updates(id, fieldValues)
	return
}

func (m *RoleService) EditRolePerm(roleId uuid.UUID, permIds []uuid.UUID) (err error) {
	err = m.repo.EditRolePerm(roleId, permIds)
	return
}

func (m *RoleService) GetByCondition(condition map[string]string) (result []model_manage.Role, total int64, err error) {
	result, total, err = m.repo.GetByCondition(condition)
	return
}
