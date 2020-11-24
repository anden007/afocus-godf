package lib

import (
	"github.com/oleiade/reflections"
	"gorm.io/gorm/schema"
)

// 通过反射转换，反向通过字段Tag信息查找字段名，然后转换成实际数据库字段名称
func GetModelDBFieldNames(module interface{}) (result map[string]string, err error) {
	result = make(map[string]string)
	err = nil

	tagMap, err := reflections.TagsDeep(module, "json")
	if err == nil {
		for fieldName, tag := range tagMap {
			if tag != "-" {
				result[tag] = schema.NamingStrategy{}.ColumnName("", fieldName)
			}
		}
	}
	return
}
