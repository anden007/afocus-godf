package interface_core

import (
	"github.com/anden007/afocus-godf/src/types"

	"github.com/kataras/iris/v12"
)

type IWxAuth interface {
	GetCurrentUser(ctx iris.Context) *types.WxAuthInfo
	CheckAuth(ctx iris.Context, returnUrl string) (result bool)
	DoAuth(ctx iris.Context)
}
