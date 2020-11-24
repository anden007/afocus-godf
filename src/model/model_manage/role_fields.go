package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"
)

func init() {
	tableName := Role{}.TableName()

	model.ModelDBFields[tableName] = make(map[string]string)

	model.ModelDBFields[tableName]["id"] = "id"

	model.ModelDBFields[tableName]["name"] = "name"

	model.ModelDBFields[tableName]["defaultRole"] = "default_role"

	model.ModelDBFields[tableName]["dataType"] = "data_type"

	model.ModelDBFields[tableName]["description"] = "description"

}