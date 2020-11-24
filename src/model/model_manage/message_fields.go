package model_manage

import (
	"github.com/anden007/afocus-godf/src/model"
)

func init() {
	tableName := Message{}.TableName()

	model.ModelDBFields[tableName] = make(map[string]string)

	model.ModelDBFields[tableName]["createTime"] = "create_time"

	model.ModelDBFields[tableName]["content"] = "content"

	model.ModelDBFields[tableName]["messageType"] = "message_type"

	model.ModelDBFields[tableName]["range"] = "range"

	model.ModelDBFields[tableName]["userIds"] = "user_ids"

	model.ModelDBFields[tableName]["senderId"] = "sender_id"

	model.ModelDBFields[tableName]["senderNickName"] = "sender_nick_name"

	model.ModelDBFields[tableName]["sendTime"] = "send_time"

	model.ModelDBFields[tableName]["id"] = "id"

	model.ModelDBFields[tableName]["title"] = "title"

}