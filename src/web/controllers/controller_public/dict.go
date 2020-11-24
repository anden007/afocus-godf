package controller_public

import (
	"github.com/anden007/afocus-godf/src/interfaces/interface_manage"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type DictController struct {
	Ctx     iris.Context
	Service interface_manage.IDictService
}

func (m *DictController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/getDataByType/{dictType}", "GetDataByType")
}

func (m *DictController) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("DictController中的Service尚未注册！")
	}
}

func (m *DictController) GetDataByType(dictType string) mvc.Result {
	success := true
	message := ""
	result, err := m.Service.GetDataByType(dictType)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message, "result": result},
	}
}
