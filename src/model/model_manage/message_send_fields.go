package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"
)

func init() {
	tableName := MessageSend{}.TableName()

	model.ModelDBFields[tableName] = make(map[string]string)

	model.ModelDBFields[tableName]["status"] = "status"

	model.ModelDBFields[tableName]["message"] = "message"

	model.ModelDBFields[tableName]["receiverId"] = "receiver_id"

	model.ModelDBFields[tableName]["receiverNickName"] = "receiver_nick_name"

	model.ModelDBFields[tableName]["readTime"] = "read_time"

	model.ModelDBFields[tableName]["createTime"] = "create_time"

	model.ModelDBFields[tableName]["id"] = "id"

	model.ModelDBFields[tableName]["messageId"] = "message_id"

}