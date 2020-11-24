package service_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/repositories/repository_manage"

	"github.com/google/uuid"
)

type PermissionService struct {
	repo *repository_manage.PermissionRepository
}

func NewPermissionService(repo *repository_manage.PermissionRepository) *PermissionService {
	instance := &PermissionService{
		repo: repo,
	}
	return instance
}

// 获取用户页面菜单数据(不包含权限节点)
func (m *PermissionService) GetMenuList(userId uuid.UUID) (result []model_manage.Permission, err error) {
	if err == nil {
		// 用户所有权限 已排序去重
		PermissionList, _ := m.repo.GetUserPermissions(userId)
		// 筛选0级页面
		for i := 0; i < len(PermissionList); i++ {
			if PermissionList[i].Level == 0 && PermissionList[i].Type == -1 {
				result = append(result, PermissionList[i])
			}
		}

		// 筛选一级页面
		firstMenuList := []*model_manage.Permission{}
		for i := 0; i < len(PermissionList); i++ {
			if PermissionList[i].Level == 1 && PermissionList[i].Type == 0 {
				firstMenuList = append(firstMenuList, &PermissionList[i])
			}
		}
		// 筛选二级页面
		secondMenuList := []*model_manage.Permission{}
		for i := 0; i < len(PermissionList); i++ {
			if PermissionList[i].Level == 2 && PermissionList[i].Type == 0 {
				secondMenuList = append(secondMenuList, &PermissionList[i])
			}
		}
		// 筛选二级页面拥有的按钮权限
		buttonPermissions := []*model_manage.Permission{}
		for i := 0; i < len(PermissionList); i++ {
			if PermissionList[i].Level == 3 && PermissionList[i].Type == 1 {
				buttonPermissions = append(buttonPermissions, &PermissionList[i])
			}
		}

		// 匹配二级页面拥有权限
		for i := 0; i < len(secondMenuList); i++ {
			for j := 0; j < len(buttonPermissions); j++ {
				if secondMenuList[i].Id == buttonPermissions[j].ParentId {
					secondMenuList[i].PermTypes = append(secondMenuList[i].PermTypes, buttonPermissions[j].ButtonType)
				}
			}
		}

		// 匹配一级页面拥有二级页面
		for i := 0; i < len(firstMenuList); i++ {
			for j := 0; j < len(secondMenuList); j++ {
				if secondMenuList[j].ParentId == firstMenuList[i].Id {
					firstMenuList[i].Children = append(firstMenuList[i].Children, secondMenuList[j])
				}
			}
		}

		// 匹配0级页面拥有一级页面
		for i := 0; i < len(result); i++ {
			for j := 0; j < len(firstMenuList); j++ {
				if firstMenuList[j].ParentId == result[i].Id {
					result[i].Children = append(result[i].Children, firstMenuList[j])
				}
			}
		}
	}
	return
}

// 获取权限列表(包含权限节点)
func (m *PermissionService) GetAll() (result []model_manage.Permission, err error) {
	PermissionList, err := m.repo.GetAll()
	if err == nil {
		// 筛选0级页面
		for i := 0; i < len(PermissionList); i++ {
			if PermissionList[i].Level == 0 && PermissionList[i].Type == -1 {
				PermissionList[i].PermTypes = []string{}
				PermissionList[i].Children = []*model_manage.Permission{}
				result = append(result, PermissionList[i])
			}
		}

		// 筛选1级页面
		var firstMenuList []*model_manage.Permission
		for i := 0; i < len(PermissionList); i++ {
			if PermissionList[i].Level == 1 && PermissionList[i].Type == 0 {
				PermissionList[i].PermTypes = []string{}
				PermissionList[i].Children = []*model_manage.Permission{}
				firstMenuList = append(firstMenuList, &PermissionList[i])
			}
		}
		// 筛选2级页面
		var secondMenuList []*model_manage.Permission
		for i := 0; i < len(PermissionList); i++ {
			if PermissionList[i].Level == 2 && PermissionList[i].Type == 0 {
				PermissionList[i].PermTypes = []string{}
				PermissionList[i].Children = []*model_manage.Permission{}
				secondMenuList = append(secondMenuList, &PermissionList[i])
			}
		}
		// 筛选3级权限
		var thirdMenuList []*model_manage.Permission
		for i := 0; i < len(PermissionList); i++ {
			if PermissionList[i].Level == 3 && PermissionList[i].Type == 1 {
				PermissionList[i].PermTypes = []string{}
				PermissionList[i].Children = []*model_manage.Permission{}
				thirdMenuList = append(thirdMenuList, &PermissionList[i])
			}
		}

		// 匹配2级页面拥有权限节点
		for i := 0; i < len(secondMenuList); i++ {
			for j := 0; j < len(thirdMenuList); j++ {
				if thirdMenuList[j].ParentId == secondMenuList[i].Id {
					secondMenuList[i].Children = append(secondMenuList[i].Children, thirdMenuList[j])
				}
			}
		}

		// 匹配1级页面拥有2级页面
		for i := 0; i < len(firstMenuList); i++ {
			for j := 0; j < len(secondMenuList); j++ {
				if secondMenuList[j].ParentId == firstMenuList[i].Id {
					firstMenuList[i].Children = append(firstMenuList[i].Children, secondMenuList[j])
				}
			}
		}

		// 匹配0级页面拥有1级页面
		for i := 0; i < len(result); i++ {
			for j := 0; j < len(firstMenuList); j++ {
				if firstMenuList[j].ParentId == result[i].Id {
					result[i].Children = append(result[i].Children, firstMenuList[j])
				}
			}
		}
	}
	return
}

func (m *PermissionService) Add(entity model_manage.Permission) (err error) {
	err = m.repo.Add(entity)
	return
}

func (m *PermissionService) Edit(entity model_manage.Permission) (err error) {
	err = m.repo.Edit(entity)
	return
}

func (m *PermissionService) DelByIds(ids []uuid.UUID) (err error) {
	err = m.repo.DelByIds(ids)
	return
}
