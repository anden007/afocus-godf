// Code generated by hero.
// source: D:\GoProjects\github.com/anden007/afocus-godf\Backend\src\generate\template\viewmodel.tpl
// DO NOT EDIT!
package template

import "bytes"

func ViewModel(pkgCode string, model string, buffer *bytes.Buffer) {
	buffer.WriteString(`
package view_model_`)
	buffer.WriteString(pkgCode)
	buffer.WriteString(`

import (
	"github.com/anden007/afocus-godf/src/model/model_`)
	buffer.WriteString(pkgCode)
	buffer.WriteString(`"
	"github.com/anden007/afocus-godf/src/web/view_model"
)

type VM_`)
	buffer.WriteString(model)
	buffer.WriteString(` struct {
	view_model.BaseViewModel
	model_`)
	buffer.WriteString(pkgCode)
	buffer.WriteString(`.`)
	buffer.WriteString(model)
	buffer.WriteString(`
}
`)

}