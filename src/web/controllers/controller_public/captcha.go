package controller_public

import (
	"strconv"

	"github.com/anden007/afocus-godf/src/interfaces/interface_core"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type CaptchaController struct {
	Ctx     iris.Context
	Captcha interface_core.ICaptcha
}

type CaptchaResult struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	CaptchaId string `json:"captchaId"`
	Image     string `json:"image"`
}

func (m *CaptchaController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/init", "Init")
	b.Handle("GET", "/reload/{id:string}", "Reload")
}

func (m *CaptchaController) AfterActivation(a mvc.AfterActivation) {
	if m.Captcha == nil {
		panic("CaptchaController中的Captcha尚未注册！")
	}
}

// @Description Init Captcha
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /api/public/captcha/init [get]
func (m *CaptchaController) Init() mvc.Result {
	reload := false
	reload, _ = strconv.ParseBool(m.Ctx.URLParam("reload"))
	id := m.Ctx.URLParam("id")
	result := CaptchaResult{
		Success:   false,
		Message:   "未知错误",
		CaptchaId: "",
		Image:     "",
	}
	if !reload || (reload && id == "") {
		captcha, err := m.Captcha.New()
		if err == nil {
			result.Success = true
			result.Message = ""
			result.CaptchaId = captcha.ID
			result.Image = captcha.Base64Image
		}
	} else {
		captcha, err := m.Captcha.Reload(id)
		if err == nil {
			result.Success = true
			result.Message = ""
			result.CaptchaId = captcha.ID
			result.Image = captcha.Base64Image
		}
	}

	return mvc.Response{
		Object: result,
	}
}

// @Description Reload Captcha
// @Accept  json
// @Produce  json
// @Param   id   path   string  false "captcha id"
// @Success 200 {string} string	"ok"
// @Router /api/public/captcha/reload/{id} [get]
func (m *CaptchaController) Reload(id string) mvc.Result {
	result := CaptchaResult{
		Success:   false,
		Message:   "未知错误",
		CaptchaId: "",
		Image:     "",
	}
	if id == "" {
		result.Success = false
		result.Message = "参数'id'必须传入"
		result.CaptchaId = ""
		result.Image = ""
	} else {
		captcha, err := m.Captcha.Reload(id)
		if err == nil {
			result.Success = true
			result.Message = ""
			result.CaptchaId = captcha.ID
			result.Image = captcha.Base64Image
		}
	}
	return mvc.Response{
		Object: result,
	}
}
