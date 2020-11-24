package view_model_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/web/view_model"

	"github.com/google/uuid"
)

type VM_User struct {
	view_model.BaseViewModel
	model_manage.User
	SelectRoles     []uuid.UUID `json:"selectRoles"`
	DepartmentTitle string      `json:"departmentTitle"`
}

func (m *VM_User) FromView() (err error) {
	// 处理页面回传的角色信息
	for _, roleId := range m.SelectRoles {
		m.User.Roles = append(m.User.Roles, &model_manage.Role{Id: roleId})
	}
	return nil
}

func (m *VM_User) FromDB() (err error) {
	if m.Department != nil {
		m.DepartmentTitle = m.Department.Title
	}
	// 转换数据模型关联角色信息
	if m.User.Roles != nil {
		for _, role := range m.User.Roles {
			m.SelectRoles = append(m.SelectRoles, role.Id)
		}
	} else {
		//将前端数据处理成空数组，否则就是null，前端可能出错
		m.SelectRoles = []uuid.UUID{}
	}
	return nil
}
