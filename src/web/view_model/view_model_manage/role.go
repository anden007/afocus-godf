package view_model_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/web/view_model"

	"github.com/google/uuid"
)

type VM_Role struct {
	view_model.BaseViewModel
	model_manage.Role
	RolePermissions []uuid.UUID `json:"rolePerms"`
	RoleDepartments []uuid.UUID `json:"roleDeps"`
}

func (m *VM_Role) FromView() (err error) {
	for _, id := range m.RolePermissions {
		m.Role.Permissions = append(m.Role.Permissions, &model_manage.Permission{Id: id})
	}
	for _, id := range m.RoleDepartments {
		m.Role.Departments = append(m.Role.Departments, &model_manage.Department{Id: id})
	}
	return nil
}

func (m *VM_Role) FromDB() (err error) {
	if m.Role.Permissions != nil {
		for _, perm := range m.Role.Permissions {
			m.RolePermissions = append(m.RolePermissions, perm.Id)
		}
	} else {
		m.RolePermissions = []uuid.UUID{}
	}

	if m.Role.Departments != nil {
		for _, dep := range m.Role.Departments {
			m.RoleDepartments = append(m.RoleDepartments, dep.Id)
		}
	} else {
		m.RoleDepartments = []uuid.UUID{}
	}
	return nil
}

type VM_Role_EditRoleDep struct {
	view_model.BaseViewModel
	RoleId   uuid.UUID   `json:"roleId"`
	DataType int         `json:"dataType"`
	RowDeps  []uuid.UUID `json:"roleDeps"`
}

type VM_Role_EditDefault struct {
	view_model.BaseViewModel
	RoleId    uuid.UUID `json:"id"`
	IsDefault bool      `json:"isDefault"`
}

type VM_Role_EditRolePerms struct {
	view_model.BaseViewModel
	RoleId    uuid.UUID   `json:"roleId"`
	RolePerms []uuid.UUID `json:"rolePerms"`
}
