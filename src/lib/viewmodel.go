package lib

import (
	"github.com/anden007/afocus-godf/src/interfaces/interface_core"

	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris/v12"
)

func ReadBody(ctx iris.Context, jsonEncoder jsoniter.API, vModel interface_core.IViewModel) (err error) {
	if ctx != nil {
		payload, err := ctx.GetBody()
		if err == nil {
			err = jsonEncoder.UnmarshalFromString(string(payload), &vModel)
			if err == nil {
				err = vModel.FromView()
			}
		}
	}
	return
}
