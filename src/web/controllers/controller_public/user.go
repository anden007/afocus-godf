package controller_public

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/interfaces/interface_manage"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	Ctx     iris.Context
	Service interface_manage.IUserService
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	jwtMiddleware := interfaces.DI().GetJwt().GetMiddleware()
	b.Handle("GET", "/info", "GetInfo", jwtMiddleware)
	b.Handle("POST", "/unlock", "UnLock")
}

func (m *UserController) AfterActivation(a mvc.AfterActivation) {
	if m.Service == nil {
		panic("UserController中的Service尚未注册！")
	}
}

func (m *UserController) GetInfo() mvc.Result {
	success := true
	message := ""
	result, err := m.Service.GetUserInfoFromJWT(m.Ctx)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message, "result": result},
	}
}

func (m *UserController) PostLogin() mvc.Result {
	result, _ := m.Service.Login(m.Ctx)
	return mvc.Response{
		Object: result,
	}
}

func (m *UserController) UnLock() mvc.Result {
	success := true
	message := ""
	err := m.Service.UnLock(m.Ctx)
	if err != nil {
		success = false
		message = err.Error()
	}
	return mvc.Response{
		Object: iris.Map{"success": success, "message": message},
	}
}

// func (m *UserController) GetTest() mvc.Result {
// 	buffer := new(bytes.Buffer)
// 	fieldList := []generator.Field{
// 		{Field: "status", Name: "状态", Level: "0", TableShow: false, SortOrder: 0.0, Searchable: true, Editable: true, Type: "text", Validate: true, SearchType: "", SearchLevel: "", Sortable: false, DefaultSort: false, DefaultSortType: ""},
// 	}
// 	options := template.AddOption{
// 		RowNum:          1,
// 		ItemWidth:       "",
// 		Width:           "100%",
// 		WangEditorWidth: "100%",
// 		QuillWidth:      "100%",
// 		ApiName:         "seller",
// 		Api:             true,
// 		Upload:          false,
// 		UploadThumb:     false,
// 		WangEditor:      false,
// 		Quill:           false,
// 		Password:        false,
// 	}
// 	template.Add(fieldList, options, buffer)
// 	return mvc.Response{
// 		Text: buffer.String(),
// 	}
// }
