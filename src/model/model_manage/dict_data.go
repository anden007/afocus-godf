package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type DictData struct {
	model.DBModel
	Id          uuid.UUID `gorm:"size:36;primaryKey" json:"id"`
	DictId      uuid.UUID `gorm:"size:36;index" json:"dictId"`
	Dict        *Dict     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Status      int       `json:"status"`
	Title       string    `gorm:"size:255" json:"title"`
	Value       string    `gorm:"size:255" json:"value"`
	SortOrder   float32   `gorm:"type:decimal(10,2)" json:"sortOrder"`
	Description string    `gorm:"size:1000" json:"description"`
}

func (DictData) TableName() string {
	return "t_dict_data"
}

type DictDataArray []*DictData

func (p DictDataArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p DictDataArray) Len() int           { return len(p) }
func (p DictDataArray) Less(i, j int) bool { return p[i].SortOrder < p[j].SortOrder }
