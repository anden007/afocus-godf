// Code generated by hero.
// source: D:\GoProjects\github.com/anden007/afocus-godf\Backend\src\generate\template\model.tpl
// DO NOT EDIT!
package template

import "bytes"

func Model(pkgCode string, model string, fieldTags map[string]string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package model_`)
	buffer.WriteString(pkgCode)
	buffer.WriteString(`

import (
	"github.com/anden007/afocus-godf/src/model"
)

func init() {
	tableName := `)
	buffer.WriteString(model)
	buffer.WriteString(`{}.TableName()

	model.ModelDBFields[tableName] = make(map[string]string)
`)
	for tag, dbFieldName := range fieldTags {
		buffer.WriteString(`
	model.ModelDBFields[tableName]["`)
		buffer.WriteString(tag)
		buffer.WriteString(`"] = "`)
		buffer.WriteString(dbFieldName)
		buffer.WriteString(`"
`)
	}
	buffer.WriteString(`
}`)

}
