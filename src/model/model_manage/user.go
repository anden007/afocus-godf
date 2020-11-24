package model_manage

import (
	"time"

	"github.com/anden007/afocus-godf/src/model"

	"github.com/google/uuid"
)

type User struct {
	model.DBModel
	Id           uuid.UUID     `gorm:"size:36;primaryKey" json:"id"`
	DepartmentId uuid.UUID     `gorm:"size:36;index" json:"departmentId"`
	Department   *Department   `gorm:"foreignKey:DepartmentId" json:"department,omitempty"`
	Roles        []*Role       `gorm:"many2many:t_user_roles;associationAutoCreate:false;associationAutoUpdate:false" json:"roles,omitempty"`
	Permissions  []*Permission `gorm:"-" json:"permissions,omitempty"`
	Status       int           `gorm:"default:0;comment:'状态 默认0正常 -1拉黑'" json:"status"`
	//Type         int           `gorm:"default:0;comment:'用户类型 0普通用户 1管理员'" json:"type"`
	UserName     string    `gorm:"size:100;unique" json:"userName"`
	Password     string    `gorm:"size:100" json:"password"`
	PassStrength int       `gorm:"default:0;comment:'密码强度'" json:"passStrength"`
	NickName     string    `gorm:"size:100" json:"nickName"`
	Avatar       string    `gorm:"size:500" json:"avatar"`
	Sex          string    `gorm:"size:10" json:"sex"`
	Mobile       string    `gorm:"size:20" json:"mobile"`
	WeiXin       string    `gorm:"size:100" json:"weixin"`
	QQ           string    `gorm:"size:20" json:"qq"`
	EMail        string    `gorm:"size:100" json:"email"`
	Address      string    `gorm:"size:500" json:"address"`
	Street       string    `gorm:"size:500" json:"street"`
	Description  string    `gorm:"size:100" json:"description"`
	CreateTime   time.Time `gorm:"size:100" json:"createTime"`
}

func (User) TableName() string {
	return "t_user"
}
