package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type Setting struct {
	model.DBModel
	Id    uuid.UUID `gorm:"size:36;primaryKey" json:"id"`
	Value string    `gorm:"size:1000" json:"value" json:"value"`
}

func (Setting) TableName() string {
	return "t_setting"
}
