package controller_public

import (
	"github.com/anden007/afocus-godf/src/interfaces/interface_manage"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type PermissionController struct {
	Ctx               iris.Context
	PermissionService interface_manage.IPermissionService
	UserService       interface_manage.IUserService
}

func (m *PermissionController) BeforeActivation(b mvc.BeforeActivation) {

	b.Handle("GET", "/getMenuList", "GetMenuList")
}

func (m *PermissionController) AfterActivation(a mvc.AfterActivation) {
	if m.PermissionService == nil {
		panic("PermissionController中的PermissionService尚未注册！")
	}
	if m.UserService == nil {
		panic("PermissionController中的UserService尚未注册！")
	}
}

func (m *PermissionController) GetMenuList() mvc.Result {
	success := true
	message := ""
	var result []model_manage.Permission
	user, err := m.UserService.GetUserInfoFromJWT(m.Ctx)
	if err == nil {
		userId, _ := uuid.Parse(user.Id)
		result, err = m.PermissionService.GetMenuList(userId)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message, "result": result},
	}
}
