package controller_manage

import (
	"strings"

	"github.com/anden007/afocus-godf/src/interfaces/interface_manage"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/web/view_model/view_model_manage"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type DictController struct {
	Ctx         iris.Context
	Service     interface_manage.IDictService
	JsonEncoder jsoniter.API
}

func (m *DictController) BeforeActivation(b mvc.BeforeActivation) {
	//字典
	b.Handle("GET", "/getAll", "GetAllDict")
	b.Handle("POST", "/add", "AddDict")
	b.Handle("POST", "/edit", "EditDict")
	b.Handle("DELETE", "/delByIds/{ids:string}", "DelDictByIds")
	//字典数据
	b.Handle("GET", "/data/getByCondition", "GetDictDataByCondition")
	b.Handle("POST", "/data/add", "AddDictData")
	b.Handle("POST", "/data/edit", "EditDictData")
	b.Handle("DELETE", "/data/delByIds/{ids:string}", "DelDictDataByIds")
}

func (m *DictController) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("DictController中的Service尚未注册！")
	}
}

func (m *DictController) AddDict() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Dict{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)

	//err := m.Ctx.ReadForm(&vModel)
	if err != nil {
		success = false
		message = err.Error()
	} else {
		if vModel.Dict.Id == uuid.Nil {
			vModel.Dict.Id = lib.NewGuid()
		}
		err = m.Service.AddOrEditDict(vModel.Dict)
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DictController) DelDictByIds(ids string) mvc.Result {
	success := true
	message := ""
	toDelIdArray := []uuid.UUID{}
	toDelIds := strings.Split(ids, ",")
	for _, idStr := range toDelIds {
		id, _ := uuid.Parse(idStr)
		toDelIdArray = append(toDelIdArray, id)
	}
	err := m.Service.DelDictByIds(toDelIdArray)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DictController) EditDict() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Dict{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err != nil {
		success = false
		message = err.Error()
	} else {
		err = m.Service.AddOrEditDict(vModel.Dict)
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DictController) GetAllDict() mvc.Result {
	success := true
	message := ""
	result, err := m.Service.GetAll()
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message, "result": result},
	}
}

func (m *DictController) AddDictData() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_DictData{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err != nil {
		success = false
		message = err.Error()
	} else {
		if vModel.DictData.Id == uuid.Nil {
			vModel.DictData.Id = lib.NewGuid()
		}
		err = m.Service.AddOrEditDictData(vModel.DictData)
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DictController) DelDictDataByIds(ids string) mvc.Result {
	success := true
	message := ""
	toDelIdArray := []uuid.UUID{}
	toDelIds := strings.Split(ids, ",")
	for _, idStr := range toDelIds {
		id, _ := uuid.Parse(idStr)
		toDelIdArray = append(toDelIdArray, id)
	}
	err := m.Service.DelDictDataByIds(toDelIdArray)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DictController) EditDictData() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_DictData{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err != nil {
		success = false
		message = err.Error()
	} else {
		if vModel.DictData.Id == uuid.Nil {
			vModel.DictData.Id = lib.NewGuid()
		}
		err = m.Service.AddOrEditDictData(vModel.DictData)
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DictController) GetDictDataByCondition() mvc.Result {
	success := true
	message := ""
	var total int64 = 0
	var result []model_manage.DictData
	var err error
	condition := m.Ctx.URLParams()
	result, total, err = m.Service.GetDictDataByCondition(condition)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message, "result": iris.Map{"content": result, "totalElements": total}},
	}
}
