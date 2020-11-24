package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type Permission struct {
	model.DBModel
	Id          uuid.UUID     `gorm:"size:36;primaryKey" json:"id"`
	ParentId    uuid.UUID     `gorm:"size:36" json:"parentId"`
	Status      int           `gorm:"default:0;comment:'是否启用 0启用 -1禁用'" json:"status"`
	Name        string        `gorm:"size:100" json:"name"`
	ShowAlways  bool          `json:"showAlways"`
	Level       int           `gorm:"default:0" json:"level"`
	Type        int           `gorm:"default:0;comment:'类型 -1顶部菜单 0页面 1具体操作'" json:"type"`
	Title       string        `gorm:"size:100" json:"title"`
	Path        string        `gorm:"size:500" json:"path"`
	URL         string        `gorm:"size:500" json:"url"`
	Component   string        `gorm:"size:100" json:"component"`
	Icon        string        `gorm:"size:100" json:"icon"`
	ButtonType  string        `gorm:"size:100" json:"buttonType"`
	Description string        `gorm:"size:1000" json:"description"`
	SortOrder   float32       `gorm:"type:decimal(10,2)" json:"sortOrder"`
	Children    []*Permission `gorm:"foreignKey:ParentId;associationForeignKey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"children,omitempty"`
	PermTypes   []string      `gorm:"-" json:"permTypes,omitempty"`
	Expand      bool          `gorm:"-" json:"expand"`
	Checked     bool          `gorm:"-" json:"checked"`
	Selected    bool          `gorm:"-" json:"selected"`
}

func (Permission) TableName() string {
	return "t_permission"
}
