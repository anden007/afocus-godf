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

type PermissionController struct {
	Ctx         iris.Context
	Service     interface_manage.IPermissionService
	JsonEncoder jsoniter.API
}

func (m *PermissionController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/add", "Add")
	b.Handle("POST", "/edit", "Edit")
	b.Handle("DELETE", "/delByIds/{ids:string}", "DelByIds")
	b.Handle("GET", "/getAll", "GetAll")
}

func (m *PermissionController) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("PermissionController中的Service尚未注册！")
	}
}

func (m *PermissionController) Add() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Permission{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		vModel.Permission.Id = lib.NewGuid()
		err = m.Service.Add(vModel.Permission)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *PermissionController) Edit() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Permission{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		err = m.Service.Edit(vModel.Permission)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *PermissionController) DelByIds(ids string) mvc.Result {
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

func (m *PermissionController) GetAll() mvc.Result {
	success := true
	message := ""
	dbResult, err := m.Service.GetAll()
	var result []view_model_manage.VM_Permission
	for i := 0; i < len(dbResult); i++ {
		des := view_model_manage.VM_Permission{}
		err = mergo.Merge(&des.Permission, dbResult[i], mergo.WithOverride, mergo.WithTransformers(&lib.UUIDTransformer{}))
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
