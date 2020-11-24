package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"
)

func init() {
	tableName := User{}.TableName()

	model.ModelDBFields[tableName] = make(map[string]string)

	model.ModelDBFields[tableName]["departmentId"] = "department_id"

	model.ModelDBFields[tableName]["mobile"] = "mobile"

	model.ModelDBFields[tableName]["department,omitempty"] = "department"

	model.ModelDBFields[tableName]["roles,omitempty"] = "roles"

	model.ModelDBFields[tableName]["sex"] = "sex"

	model.ModelDBFields[tableName]["qq"] = "qq"

	model.ModelDBFields[tableName]["address"] = "address"

	model.ModelDBFields[tableName]["id"] = "id"

	model.ModelDBFields[tableName]["nickName"] = "nick_name"

	model.ModelDBFields[tableName]["passStrength"] = "pass_strength"

	model.ModelDBFields[tableName]["email"] = "e_mail"

	model.ModelDBFields[tableName]["description"] = "description"

	model.ModelDBFields[tableName]["avatar"] = "avatar"

	model.ModelDBFields[tableName]["status"] = "status"

	model.ModelDBFields[tableName]["street"] = "street"

	model.ModelDBFields[tableName]["userName"] = "user_name"

	model.ModelDBFields[tableName]["weixin"] = "wei_xin"

	model.ModelDBFields[tableName]["permissions,omitempty"] = "permissions"

	model.ModelDBFields[tableName]["createTime"] = "create_time"

	model.ModelDBFields[tableName]["password"] = "password"

}
