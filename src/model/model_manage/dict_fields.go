package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"
)

func init() {
	tableName := Dict{}.TableName()

	model.ModelDBFields[tableName] = make(map[string]string)

	model.ModelDBFields[tableName]["type"] = "type"

	model.ModelDBFields[tableName]["sortOrder"] = "sort_order"

	model.ModelDBFields[tableName]["description"] = "description"

	model.ModelDBFields[tableName]["id"] = "id"

	model.ModelDBFields[tableName]["title"] = "title"

}