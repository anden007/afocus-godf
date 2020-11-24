package controller_manage

import (
	"errors"
	"strings"

	"github.com/anden007/afocus-godf/src/interfaces/interface_manage"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/web/view_model/view_model_manage"

	"github.com/google/uuid"
	"github.com/imdario/mergo"
	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type MessageController struct {
	Ctx         iris.Context
	JsonEncoder jsoniter.API
	Service     interface_manage.IMessageService
	UserService interface_manage.IUserService
}

func (m *MessageController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/add", "Add")
	b.Handle("DELETE", "/delByIds/{ids:string}", "DelByIds")
	b.Handle("POST", "/edit", "Edit")
	b.Handle("GET", "/getById/{id:string}", "GetById")
	b.Handle("GET", "/getAll", "GetAll")
	b.Handle("GET", "/getByCondition", "GetByCondition")
}

func (m *MessageController) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("MessageController中的Service尚未注册！")
	}
}

func (m *MessageController) Add() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Message{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		vModel.Message.Id = lib.NewGuid()
		if userInfo, jwtErr := m.UserService.GetUserInfoFromJWT(m.Ctx); jwtErr == nil {
			if userInfo != nil {
				if userId, uuidErr := uuid.Parse(userInfo.Id); uuidErr == nil {
					vModel.Message.SenderId = userId
					vModel.Message.SenderNickName = userInfo.NickName
					err = m.Service.Add(vModel.Message)
				} else {
					err = uuidErr
				}
			} else {
				err = errors.New("用户尚未登录")
			}
		} else {
			err = jwtErr
		}
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *MessageController) DelByIds(ids string) mvc.Result {
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

func (m *MessageController) Edit() mvc.Result {
	success := true
	message := ""
	vModel := view_model_manage.VM_Message{}
	err := lib.ReadBody(m.Ctx, m.JsonEncoder, &vModel)
	if err == nil {
		err = m.Service.Edit(vModel.Message)
	}
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

func (m *MessageController) GetById(id string) mvc.Result {
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

func (m *MessageController) GetAll() mvc.Result {
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

func (m *MessageController) GetByCondition() mvc.Result {
	success := true
	message := ""
	var total int64 = 0
	var dbResult []model_manage.Message
	var err error
	condition := m.Ctx.URLParams()
	dbResult, total, err = m.Service.GetByCondition(condition)
	var result []view_model_manage.VM_Message
	for i := 0; i < len(dbResult); i++ {
		des := view_model_manage.VM_Message{}
		err = mergo.Merge(&des.Message, dbResult[i], mergo.WithOverride, mergo.WithTransformers(&lib.UUIDTransformer{}))
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
