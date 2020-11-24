package view_model_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/web/view_model"
)

type VM_Task struct {
	view_model.BaseViewModel
	model_manage.Task
}
