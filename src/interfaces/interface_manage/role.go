package interface_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type IRoleService interface {
	Add(entity model_manage.Role) (err error)
	Edit(entity model_manage.Role) (err error)
	Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error)
	DelByIds(ids []uuid.UUID) (err error)
	EditRoleDep(roleId uuid.UUID, dataType int, depIds []uuid.UUID) (err error)
	GetAll() (result []model_manage.Role, err error)
	EditRolePerm(roleId uuid.UUID, permIds []uuid.UUID) (err error)
	GetByCondition(condition map[string]string) (result []model_manage.Role, total int64, err error)
}
