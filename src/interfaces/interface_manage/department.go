package interface_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/google/uuid"
)

type IDepartmentService interface {
	GetByParentId(parentId uuid.UUID) (result []model_manage.Department, err error)
	Add(entity model_manage.Department) (err error)
	DelByIds(ids []uuid.UUID) (err error)
	Edit(entity model_manage.Department) (err error)
}
