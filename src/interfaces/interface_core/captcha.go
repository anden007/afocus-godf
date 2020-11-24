package interface_core

import "github.com/anden007/afocus-godf/src/types"

type ICaptcha interface {
	New() (result *types.Captcha, err error)
	Reload(id string) (result *types.Captcha, err error)
	Verify(id string, digits string) (result bool, err error)
}
