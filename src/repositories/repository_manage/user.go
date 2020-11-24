package repository_manage

import (
	"errors"
	"time"

	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	instance := new(UserRepository)
	return instance
}

func (m *UserRepository) Login(userName string, password string) (user model_manage.User, err error) {
	db := interfaces.DI().GetDataBase()
	md5Password := lib.Md5Hash(password)
	err = db.GetDB().Model(model_manage.User{}).Where(&model_manage.User{UserName: userName, Password: md5Password}).Preload("Roles").First(&user).Error
	return
}

func (m *UserRepository) GetByDepartmentId(departmentId uuid.UUID) (result []model_manage.User, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(model_manage.User{}).Find(&result, &model_manage.User{DepartmentId: departmentId}).Error
	return
}

func (m *UserRepository) GetByCondition(condition map[string]string) (result []model_manage.User, total int64, err error) {
	db := interfaces.DI().GetDataBase()
	Query := db.GetDB().Model(model_manage.User{})
	CountQuery := db.GetDB().Model(model_manage.User{})
	err = lib.NewQueryCondition().GetQuery(model_manage.User{}.TableName(), condition, Query, CountQuery)
	if err == nil {
		err = CountQuery.Count(&total).Error
		err = Query.Preload("Roles").Preload("Department").Find(&result).Error
	}
	return
}

func (m *UserRepository) Add(entity model_manage.User) (err error) {
	db := interfaces.DI().GetDataBase()
	entity.Password, err = m.CreatePassword(entity.Password)
	if err == nil {
		if entity.CreateTime.IsZero() {
			entity.CreateTime = time.Now()
		}
		err = db.GetDB().Create(&entity).Error
	}
	return
}

func (m *UserRepository) Edit(entity model_manage.User) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Save(&entity).Error
	return
}

func (m *UserRepository) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.User{}).Where(model_manage.User{Id: id}).Updates(fieldValues).Error
	return
}

func (m *UserRepository) DelByIds(ids []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Delete(model_manage.User{}, "id in (?)", ids).Error
	return
}

func (m *UserRepository) GetByIds(ids []uuid.UUID) (result []model_manage.User, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(model_manage.User{}).Where("id in (?)", ids).Find(&result).Error
	return
}

func (m *UserRepository) ResetPassword(ids []uuid.UUID) (err error) {
	md5Password, _ := m.CreatePassword("123456")
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.User{}).Where("id in (?)", ids).Update("Password", md5Password).Error
	return
}

func (m *UserRepository) ModifyPass(id uuid.UUID, oldPassword string, newPassword string) (err error) {
	if len(newPassword) < 6 {
		err = errors.New("密码不满足系统要求")
	} else {
		md5OldPassword, _ := m.CreatePassword(oldPassword)
		md5NewPassword, _ := m.CreatePassword(newPassword)
		db := interfaces.DI().GetDataBase()
		result := model_manage.User{}
		err = db.GetDB().Model(&model_manage.User{}).Where(&model_manage.User{Id: id, Password: md5OldPassword}).First(&result).Error
		if err == nil {
			err = db.GetDB().Model(&model_manage.User{}).Where("id in (?)", id).Update("Password", md5NewPassword).Error
		}
	}
	return
}

func (m *UserRepository) CreatePassword(password string) (result string, err error) {
	err = nil
	if len(password) < 6 {
		err = errors.New("密码不满足系统要求")
	} else {
		result = lib.Md5Hash(password)
	}
	return
}

// 编辑用户角色
func (m *UserRepository) EditUserRole(userId uuid.UUID, roleIds []uuid.UUID) (err error) {
	db := interfaces.DI().GetDataBase()
	var newRoles []model_manage.Role
	err = db.GetDB().Model(&model_manage.Role{}).Where("id in (?)", roleIds).Find(&newRoles).Error
	if err == nil {
		var user model_manage.User
		err = db.GetDB().First(&user, "id = ?", userId).Error
		err = db.GetDB().Model(&user).Association("Roles").Replace(&newRoles)
	}
	return
}

func (m *UserRepository) GetAll() (result []model_manage.User, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.User{}).Find(&result).Error
	return
}
