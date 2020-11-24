package interface_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type IPermissionService interface {
	GetAll() (result []model_manage.Permission, err error)
	Add(entity model_manage.Permission) (err error)
	DelByIds(ids []uuid.UUID) (err error)
	Edit(entity model_manage.Permission) (err error)
	GetMenuList(userId uuid.UUID) (result []model_manage.Permission, err error)
}
