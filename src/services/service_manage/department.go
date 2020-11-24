package service_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/repositories/repository_manage"

	"github.com/google/uuid"
)

type DepartmentService struct {
	repo *repository_manage.DepartmentRepository
}

func NewDepartmentService(repo *repository_manage.DepartmentRepository) *DepartmentService {
	instance := &DepartmentService{
		repo: repo,
	}
	return instance
}

func (m *DepartmentService) GetByParentId(pid uuid.UUID) (result []model_manage.Department, err error) {
	result, err = m.repo.GetByParentId(pid)
	return
}

func (m *DepartmentService) Add(entity model_manage.Department) (err error) {
	err = m.repo.Add(entity)
	return
}

func (m *DepartmentService) Edit(entity model_manage.Department) (err error) {
	err = m.repo.Edit(entity)
	return
}

func (m *DepartmentService) DelByIds(ids []uuid.UUID) (err error) {
	err = m.repo.DelByIds(ids)
	return
}
