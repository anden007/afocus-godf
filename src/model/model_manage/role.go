package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type Role struct {
	model.DBModel
	Id          uuid.UUID     `gorm:"size:36;primaryKey" json:"id"`
	Name        string        `gorm:"size:100;comment:'角色名 以ROLE_开头'" json:"name"`
	DefaultRole bool          `gorm:"default:0;comment:'是否为注册默认角色'" json:"defaultRole"`
	DataType    int           `gorm:"default:0;comment:'数据权限类型 0全部默认 1自定义 2本部门及以下 3本部门 4仅本人'" json:"dataType"`
	Description string        `gorm:"size:100" json:"description"`
	Permissions []*Permission `gorm:"many2many:t_role_permissions;associationAutocCeate:false;associationAutoUpdate:false;" json:"-"`
	Departments []*Department `gorm:"many2many:t_role_departments;associationAutocReate:false;associationAutoUpdate:false;" json:"-"`
}

func (Role) TableName() string {
	return "t_role"
}
