package interface_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/types"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type IUserService interface {
	GetByCondition(condition map[string]string) (result []model_manage.User, total int64, err error)
	Add(entity model_manage.User) (err error)
	Edit(entity model_manage.User) (err error)
	Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error)
	DelByIds(ids []uuid.UUID) (err error)
	GetAll() (result []model_manage.User, err error)
	GetByIds(ids []uuid.UUID) (result []model_manage.User, err error)
	Login(ctx iris.Context) (result iris.Map, err error)
	UnLock(ctx iris.Context) (err error)
	GetUserInfoFromJWT(ctx iris.Context) (result *types.BaseUserInfo, err error)
	GetByDepartmentId(departmentId uuid.UUID) (result []model_manage.User, err error)
	ResetPassword(ids []uuid.UUID) (err error)
	ModifyPass(id uuid.UUID, oldPassword string, newPassword string) (err error)
	CreatePassword(password string) (result string, err error)
	EditUserRole(userId uuid.UUID, roleIds []uuid.UUID) (err error)
}
