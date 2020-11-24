package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"
)

func init() {
	tableName := Task{}.TableName()

	model.ModelDBFields[tableName] = make(map[string]string)

	model.ModelDBFields[tableName]["lastRunTime"] = "last_run_time"

	model.ModelDBFields[tableName]["id"] = "id"

	model.ModelDBFields[tableName]["funcName"] = "func_name"

	model.ModelDBFields[tableName]["success"] = "success"

	model.ModelDBFields[tableName]["result"] = "result"

	model.ModelDBFields[tableName]["creatorId"] = "creator_id"

	model.ModelDBFields[tableName]["createTime"] = "create_time"

}