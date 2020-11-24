package service_manage

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/interfaces/interface_core"
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/repositories/repository_manage"
	"github.com/anden007/afocus-godf/src/types"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type UserService struct {
	repo    *repository_manage.UserRepository
	captcha interface_core.ICaptcha
	jwt     interface_core.IJwtService
}

func NewUserService(repo *repository_manage.UserRepository) *UserService {
	instance := &UserService{
		repo:    repo,
		captcha: interfaces.DI().GetCaptcha(),
		jwt:     interfaces.DI().GetJwt(),
	}
	return instance
}

func (m *UserService) Login(ctx iris.Context) (result iris.Map, err error) {
	loginCode := ctx.FormValue("username")
	password := ctx.FormValue("password")
	vCode := ctx.FormValue("code")
	captchaID := ctx.FormValue("captchaId")
	// saveLogin, err := strconv.ParseBool(ctx.FormValue("saveLogin"))
	if check, _ := m.captcha.Verify(captchaID, vCode); check {
		user, err := m.repo.Login(loginCode, password)
		if err == nil {
			//var Menus []string = []string{}
			//var Permissions []string
			var Roles []string = []string{}
			if user.Roles != nil {
				for _, role := range user.Roles {
					Roles = append(Roles, role.Name)
				}
			}
			claims := types.BaseUserInfo{
				Id:       user.Id.String(),
				Roles:    Roles,
				NickName: user.NickName,
				Avatar:   user.Avatar,
				UserName: user.UserName,
				Sex:      user.Sex,
				Mobile:   user.Mobile,
				WeiXin:   user.WeiXin,
				QQ:       user.QQ,
				EMail:    user.EMail,
				Address:  user.Address,
				Street:   user.Street,
			}
			token, _ := m.jwt.CreateToken(ctx, claims)
			return iris.Map{"success": true, "message": "", "accessToken": token}, err
		}
		return iris.Map{"success": false, "message": "登陆失败"}, err
	}
	return iris.Map{"success": false, "message": "验证码错误"}, err
}
func (m *UserService) UnLock(ctx iris.Context) (err error) {
	password := ctx.FormValue("password")
	if claims := m.jwt.GetClaims(ctx); claims != nil {
		userInfoClaims := claims.(types.BaseUserInfoClaims)
		userInfo := userInfoClaims.BaseUserInfo
		_, err = m.repo.Login(userInfo.UserName, password)
	}
	return
}

func (m *UserService) GetUserInfoFromJWT(ctx iris.Context) (result *types.BaseUserInfo, err error) {
	if claims := m.jwt.GetClaims(ctx); claims != nil {
		userInfoClaims := claims.(*types.BaseUserInfoClaims)
		result = &userInfoClaims.BaseUserInfo
	}
	return
}

func (m *UserService) GetByDepartmentId(departmentId uuid.UUID) (result []model_manage.User, err error) {
	result, err = m.repo.GetByDepartmentId(departmentId)
	return
}

func (m *UserService) GetByCondition(condition map[string]string) (result []model_manage.User, total int64, err error) {
	result, total, err = m.repo.GetByCondition(condition)
	return
}

func (m *UserService) Add(entity model_manage.User) (err error) {
	err = m.repo.Add(entity)
	return
}

func (m *UserService) Edit(entity model_manage.User) (err error) {
	err = m.repo.Edit(entity)
	return
}
func (m *UserService) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	err = m.repo.Updates(id, fieldValues)
	return
}

func (m *UserService) DelByIds(ids []uuid.UUID) (err error) {
	err = m.repo.DelByIds(ids)
	return
}

func (m *UserService) GetByIds(ids []uuid.UUID) (result []model_manage.User, err error) {
	result, err = m.repo.GetByIds(ids)
	return
}

func (m *UserService) ResetPassword(ids []uuid.UUID) (err error) {
	err = m.repo.ResetPassword(ids)
	return
}

func (m *UserService) ModifyPass(id uuid.UUID, oldPassword string, newPassword string) (err error) {
	err = m.repo.ModifyPass(id, oldPassword, newPassword)
	return
}

func (m *UserService) CreatePassword(password string) (result string, err error) {
	result, err = m.repo.CreatePassword(password)
	return
}
func (m *UserService) EditUserRole(userId uuid.UUID, roleIds []uuid.UUID) (err error) {
	return m.repo.EditUserRole(userId, roleIds)
}
func (m *UserService) GetAll() (result []model_manage.User, err error) {
	result, err = m.repo.GetAll()
	return
}
