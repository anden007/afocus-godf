package view_model_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/web/view_model"
)

type VM_Department struct {
	view_model.BaseViewModel
	model_manage.Department
	ParentTitle string   `json:"parentTitle"`
	MainHeader  []string `json:"mainHeader,omitempty"`
	ViceHeader  []string `json:"viceHeader,omitempty"`
}

func (m *VM_Department) FromDB() (err error) {
	if m.Parent != nil {
		m.ParentTitle = m.Parent.Title
	}
	return nil
}
