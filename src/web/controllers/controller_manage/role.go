package controller_manage

import (
	"strings"

	"github.com/anden007/afocus-godf/src/interfaces/interface_manage"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/web/view_model/view_model_manage"

	"github.com/google/uuid"
	"github.com/imdario/mergo"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type RoleController struct {
	Ctx         iris.Context
	Service     interface_manage.IRoleService
	JsonEncoder jsoniter.API
}

func (m *RoleController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/add", "Add")
	b.Handle("DELETE", "/delByIds/{ids:string}", "DelByIds")
	b.Handle("POST", "/edit", "Edit")
	b.Handle("GET", "/getAll", "GetAll")
	b.Handle("GET", "/getAllByPage", "GetAllByPage")

	b.Handle("POST", "/editRoleDep", "EditRoleDep")
	b.Handle("POST", "/setDefault", "EditDefault")
	b.Handle("POST", "/editRolePerm", "EditRolePerm")
}

func (m *RoleController) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("RoleController中的Service尚未注册！")
	}
}

func (m *RoleController) Add() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Role{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		vModel.Role.Id = lib.NewGuid()
		err = m.Service.Add(vModel.Role)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *RoleController) DelByIds(ids string) mvc.Result {
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

func (m *RoleController) Edit() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Role{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		err = m.Service.Edit(vModel.Role)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *RoleController) GetAll() mvc.Result {
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

func (m *RoleController) GetAllByPage() mvc.Result {
	success := true
	message := ""
	var total int64 = 0
	var dbResult []model_manage.Role
	var err error
	condition := m.Ctx.URLParams()
	dbResult, total, err = m.Service.GetByCondition(condition)
	var result []view_model_manage.VM_Role
	for i := 0; i < len(dbResult); i++ {
		des := view_model_manage.VM_Role{}
		err = mergo.Merge(&des.Role, dbResult[i], mergo.WithOverride, mergo.WithTransformers(&lib.UUIDTransformer{}))
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
		Object: iris.Map{"success": success, "message": message, "result": iris.Map{"content": result, "totalElements": total}},
	}
}

func (m *RoleController) EditRoleDep() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Role_EditRoleDep{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		err = m.Service.EditRoleDep(vModel.RoleId, vModel.DataType, vModel.RowDeps)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *RoleController) EditDefault() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Role_EditDefault{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		err = m.Service.Updates(vModel.RoleId, iris.Map{"DefaultRole": vModel.IsDefault})
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *RoleController) EditRolePerm() mvc.Result {
	success := true
	message := ""
	var err error = nil
	editRolePerms := view_model_manage.VM_Role_EditRolePerms{}
	if err = lib.ReadBody(m.Ctx, m.JsonEncoder, &editRolePerms); err == nil {
		err = m.Service.EditRolePerm(editRolePerms.RoleId, editRolePerms.RolePerms)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}
