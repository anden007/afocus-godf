package service_core

import (
	"github.com/anden007/afocus-godf/src/interfaces/interface_core"
	"github.com/anden007/afocus-godf/src/types"

	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	captcha *base64Captcha.Captcha
}

func NewCaptcha(cache interface_core.ICache) *Captcha {
	instance := new(Captcha)
	store := NewRedisStore(cache)
	driver := base64Captcha.NewDriverDigit(32, 110, 4, 0.1, 10)
	instance.captcha = base64Captcha.NewCaptcha(driver, store)
	return instance
}

func (m *Captcha) New() (result *types.Captcha, err error) {
	err = nil
	id, digits, answer := m.captcha.Driver.GenerateIdQuestionAnswer()
	item, err := m.captcha.Driver.DrawCaptcha(digits)
	if err == nil {
		result = &types.Captcha{
			ID:          id,
			Digits:      answer,
			Base64Image: item.EncodeB64string(),
		}
		m.captcha.Store.Set(result.ID, result.Digits)
	}
	return result, err
}

func (m *Captcha) Reload(id string) (result *types.Captcha, err error) {
	err = nil
	_, digits, answer := m.captcha.Driver.GenerateIdQuestionAnswer()
	item, err := m.captcha.Driver.DrawCaptcha(digits)
	if err == nil {
		result = &types.Captcha{
			ID:          id,
			Digits:      answer,
			Base64Image: item.EncodeB64string(),
		}
		m.captcha.Store.Set(result.ID, result.Digits)
	}
	return result, err
}

func (m *Captcha) Verify(id string, digits string) (result bool, err error) {
	result = m.captcha.Store.Verify(id, digits, true)
	return result, err
}
