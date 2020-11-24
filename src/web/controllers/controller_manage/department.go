package controller_manage

import (
	"strings"

	"github.com/anden007/afocus-godf/src/interfaces/interface_manage"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/web/view_model/view_model_manage"

	"github.com/google/uuid"
	"github.com/imdario/mergo"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type DepartmentController struct {
	Ctx         iris.Context
	Service     interface_manage.IDepartmentService
	JsonEncoder jsoniter.API
}

func (m *DepartmentController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/add", "Add")
	b.Handle("DELETE", "/delByIds/{ids:string}", "DelByIdsHandler")
	b.Handle("POST", "/edit", "Edit")
	b.Handle("GET", "/getByParentId/{pid}", "GetByParentIdHandler")
}

func (m *DepartmentController) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("DepartmentController中的Service尚未注册！")
	}
}

func (m *DepartmentController) Add() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Department{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err != nil {
		success = false
		message = err.Error()
	} else {
		vModel.Department.Id = lib.NewGuid()
		err = m.Service.Add(vModel.Department)
		if err != nil {
			success = false
			message = err.Error()
		}
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DepartmentController) Edit() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Department{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err != nil {
		success = false
		message = err.Error()
	} else {
		err = m.Service.Edit(vModel.Department)
		if err != nil {
			success = false
			message = err.Error()
		}
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DepartmentController) DelByIdsHandler(ids string) mvc.Result {
	success := true
	message := ""
	toDelIdArray := []uuid.UUID{}
	toDelIds := strings.Split(ids, ",")
	for _, idStr := range toDelIds {
		id, _ := uuid.Parse(idStr)
		toDelIdArray = append(toDelIdArray, id)
	}
	err := m.Service.DelByIds(toDelIdArray)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *DepartmentController) GetByParentIdHandler(pid string) mvc.Result {
	success := true
	message := ""
	PId, _ := uuid.Parse(pid)
	dbResult, err := m.Service.GetByParentId(PId)
	var result []view_model_manage.VM_Department
	for i := 0; i < len(dbResult); i++ {
		des := view_model_manage.VM_Department{}
		err = mergo.Merge(&des.Department, dbResult[i], mergo.WithOverride, mergo.WithTransformers(&lib.UUIDTransformer{}))
		if err == nil {
			err = des.FromDB()
		}
		result = append(result, des)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message, "result": result},
	}
}
