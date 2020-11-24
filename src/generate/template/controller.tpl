<%: func Controller(pkgCode string, model string, buffer *bytes.Buffer) %>
package controller_<%==s pkgCode%>

import (
	"github.com/anden007/afocus-godf/src/interfaces/interface_<%==s pkgCode%>"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_<%==s pkgCode%>"
	"github.com/anden007/afocus-godf/src/web/view_model/view_model_<%==s pkgCode%>"
	"github.com/imdario/mergo"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/google/uuid"
	"strings"
)

type <%==s model%>Controller struct {
	Ctx         iris.Context
	JsonEncoder jsoniter.API
	Service     interface_<%==s pkgCode%>.I<%==s model%>Service
}

func (m *<%==s model%>Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/add", "Add")
	b.Handle("DELETE", "/delByIds/{ids:string}", "DelByIds")
	b.Handle("POST", "/edit", "Edit")
	b.Handle("GET", "/getById/{id:string}", "GetById")
	b.Handle("GET", "/getAll", "GetAll")
	b.Handle("GET", "/getByCondition", "GetByCondition")
}

func (m *<%==s model%>Controller) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("<%==s model%>Controller中的Service尚未注册！")
	}
}

func (m *<%==s model%>Controller) Add() mvc.Result {
	success := true
	message := ""
	vModel := view_model_<%==s pkgCode%>.VM_<%==s model%>{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		vModel.<%==s model%>.Id = lib.NewGuid()
		err = m.Service.Add(vModel.<%==s model%>)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *<%==s model%>Controller) DelByIds(ids string) mvc.Result {
	success := true
	message := ""
	toDelIdArray := []uuid.UUID{}
	toDelIds := strings.Split(ids, ",")
	for _, idStr := range toDelIds {
		id, _ := uuid.Parse(idStr)
		toDelIdArray = append(toDelIdArray, id)
	}
	err := m.Service.DelByIds(toDelIdArray)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *<%==s model%>Controller) Edit() mvc.Result {
	success := true
	message := ""
	vModel := view_model_<%==s pkgCode%>.VM_<%==s model%>{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		err = m.Service.Edit(vModel.<%==s model%>)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *<%==s model%>Controller) GetById(id string) mvc.Result {
	success := true
	message := ""
	Id, _ := uuid.Parse(id)
	result, err := m.Service.GetById(Id)
	if err != nil {
		success = false
		message = err.Error()
	}
	jsonStr, _ := m.JsonEncoder.Marshal(iris.Map{"success": success, "message": message, "result": result})
	return mvc.Response{
		ContentType: "application/json;charset=UTF-8",
		Text:        string(jsonStr),
	}
}

func (m *<%==s model%>Controller) GetAll() mvc.Result {
	success := true
	message := ""
	result, err := m.Service.GetAll()
	if err != nil {
		success = false
		message = err.Error()
	}
	jsonStr, _ := m.JsonEncoder.Marshal(iris.Map{"success": success, "message": message, "result": result})
	return mvc.Response{
		ContentType: "application/json;charset=UTF-8",
		Text:        string(jsonStr),
	}
}

func (m *<%==s model%>Controller) GetByCondition() mvc.Result {
	success := true
	message := ""
	var total int64 = 0
	var dbResult []model_<%==s pkgCode%>.<%==s model%>
	var err error
	condition := m.Ctx.URLParams()
	dbResult, total, err = m.Service.GetByCondition(condition)
	var result []view_model_<%==s pkgCode%>.VM_<%==s model%>
	for i := 0; i < len(dbResult); i++ {
		des := view_model_<%==s pkgCode%>.VM_<%==s model%>{}
		err = mergo.Merge(&des.<%==s model%>, dbResult[i], mergo.WithOverride, mergo.WithTransformers(&lib.UUIDTransformer{}))
		if err == nil {
			err = des.FromDB()
		}
		result = append(result, des)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	jsonStr, _ := m.JsonEncoder.MarshalToString(iris.Map{"success": success, "message": message, "result": iris.Map{"content": result, "totalElements": total}})
	return mvc.Response{
		ContentType: "application/json;charset=UTF-8",
		Text:        jsonStr,
	}
}
