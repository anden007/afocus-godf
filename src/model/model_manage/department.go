package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type Department struct {
	model.DBModel
	Id        uuid.UUID   `gorm:"size:36;primaryKey" json:"id"`
	ParentId  uuid.UUID   `gorm:"size:36;index" json:"parentId"`
	Parent    *Department `gorm:"foreignKey:ParentId;references:Id;constraint:OnUpdate:SET NULL,OnDelete:SET NULL" json:"-"`
	Users     []*User     `gorm:"foreignKey:DepartmentId" json:"users,omitempty"`
	Status    int         `gorm:"default:0" json:"status"`
	IsParent  bool        `gorm:"default:0;comment:'是否为父节点(含子节点) 默认false'" json:"isParent"`
	Title     string      `gorm:"size:100" json:"title"`
	SortOrder float32     `gorm:"type:decimal(10,2)" json:"sortOrder"`
}

func (Department) TableName() string {
	return "t_department"
}
