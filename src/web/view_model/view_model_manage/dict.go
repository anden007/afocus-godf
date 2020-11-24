package view_model_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/web/view_model"
)

type VM_Dict struct {
	view_model.BaseViewModel
	model_manage.Dict
}

type VM_DictData struct {
	view_model.BaseViewModel
	model_manage.DictData
}
