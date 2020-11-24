package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type Dict struct {
	model.DBModel
	Id          uuid.UUID   `gorm:"size:36;primaryKey" json:"id"`
	Data        []*DictData `gorm:"foreignKey:DictId" json:"-"`
	Title       string      `gorm:"size:100" json:"title"`
	Type        string      `gorm:"size:100" json:"type"`
	SortOrder   float32     `gorm:"type:decimal(10,2)" json:"sortOrder"`
	Description string      `gorm:"size:1000" json:"description"`
}

func (Dict) TableName() string {
	return "t_dict"
}
