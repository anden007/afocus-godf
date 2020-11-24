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

type UserController struct {
	Ctx         iris.Context
	Service     interface_manage.IUserService
	JsonEncoder jsoniter.API
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/add", "Add")
	b.Handle("DELETE", "/delByIds/{ids:string}", "DelByIds")
	b.Handle("POST", "/edit", "Edit")
	b.Handle("GET", "/getByCondition", "GetByCondition")

	b.Handle("GET", "/getByDepartmentId/{departmentId:string}", "GetByDepartmentId")
	b.Handle("POST", "/resetPass", "ResetPassword")
	b.Handle("POST", "/modifyPass", "ModifyPass")
	b.Handle("POST", "/disable/{uid:string}", "Disable")
	b.Handle("POST", "/enable/{uid:string}", "Enable")
}

func (m *UserController) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("UserController中的Service尚未注册！")
	}
}

func (m *UserController) Add() mvc.Result {
	success := true
	message := ""

	vModel := view_model_manage.VM_User{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err != nil {
		success = false
		message = err.Error()
	} else {
		vModel.User.Id = lib.NewGuid()
		err = m.Service.Add(vModel.User)
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *UserController) DelByIds(ids string) mvc.Result {
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

func (m *UserController) Edit() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_User{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err != nil {
		success = false
		message = err.Error()
	} else {
		if vModel.Roles != nil {
			err = m.Service.EditUserRole(vModel.Id, vModel.SelectRoles)
		}
		if err == nil {
			err = m.Service.Updates(vModel.User.Id, map[string]interface{}{
				"UserName":     vModel.UserName,
				"EMail":        vModel.EMail,
				"Mobile":       vModel.Mobile,
				"Sex":          vModel.Sex,
				"Avatar":       vModel.Avatar,
				"DepartmentId": vModel.DepartmentId,
			})
		}
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *UserController) GetByCondition() mvc.Result {
	success := true
	message := ""
	var total int64 = 0
	var dbResult []model_manage.User
	var err error
	condition := m.Ctx.URLParams()
	dbResult, total, err = m.Service.GetByCondition(condition)
	// 将数据库结果转换成VM模型
	// 由于go没有setter和getter，所以采用合并方式将db对象转换为vm对象，这种方式很灵活
	var result []view_model_manage.VM_User
	for i := 0; i < len(dbResult); i++ {
		des := view_model_manage.VM_User{}
		err = mergo.Merge(&des.User, dbResult[i], mergo.WithOverride, mergo.WithTransformers(&lib.UUIDTransformer{}))
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

func (m *UserController) ResetPassword() mvc.Result {
	success := true
	message := ""
	ids := m.Ctx.FormValue("ids")
	toResetIds := strings.Split(ids, ",")
	toResetIdArray := []uuid.UUID{}
	for _, idStr := range toResetIds {
		id, _ := uuid.Parse(idStr)
		toResetIdArray = append(toResetIdArray, id)
	}
	err := m.Service.ResetPassword(toResetIdArray)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *UserController) ModifyPass() mvc.Result {
	success := true
	message := ""
	uid, _ := uuid.Parse(m.Ctx.FormValue("id"))
	oldPassword := m.Ctx.FormValue("oldPassword")
	newPassword := m.Ctx.FormValue("newPassword")
	err := m.Service.ModifyPass(uid, oldPassword, newPassword)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *UserController) GetByDepartmentId(departmentId string) mvc.Result {
	success := true
	message := ""
	result, err := m.Service.GetByDepartmentId(uuid.Must(uuid.Parse(departmentId)))
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message, "result": result},
	}
}

func (m *UserController) Disable(userId string) mvc.Result {
	success := true
	message := ""
	err := m.Service.Updates(uuid.Must(uuid.Parse(userId)), iris.Map{"Status": -1})
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *UserController) Enable(userId string) mvc.Result {
	success := true
	message := ""
	err := m.Service.Updates(uuid.Must(uuid.Parse(userId)), iris.Map{"Status": 0})
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}
