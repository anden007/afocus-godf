package controller_public

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/anden007/afocus-godf/src/lib"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UploadController struct {
	Ctx iris.Context
}

func (m *UploadController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/file", "UploadFile")
}

func (m *UploadController) UploadFile() mvc.Result {
	fileUrl := "#"
	resultUrl := "#"
	// Get the file from the request.
	file, info, err := m.Ctx.FormFile("file")
	if err != nil {
		return mvc.Response{
			Code:   iris.StatusInternalServerError,
			Object: iris.Map{"success": false, "message": err.Error()},
		}
	}

	defer file.Close()
	fPath := "./uploads/" + time.Now().Format("2006-01")
	fName := lib.NewGuidString() + path.Ext(info.Filename)
	fileUrl = fmt.Sprintf("%s/%s", fPath, fName)
	resultUrl = fmt.Sprintf("%s/%s", "/uploads/"+time.Now().Format("2006-01"), fName)
	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	err = os.MkdirAll(fPath, os.ModePerm)
	if err != nil {
		return mvc.Response{
			Code:   iris.StatusInternalServerError,
			Object: iris.Map{"success": false, "message": err.Error()},
		}
	}
	out, err := os.OpenFile(fileUrl,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return mvc.Response{
			Code:   iris.StatusInternalServerError,
			Object: iris.Map{"success": false, "message": err.Error()},
		}
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return mvc.Response{
			Code:   iris.StatusInternalServerError,
			Object: iris.Map{"success": false, "message": err.Error()},
		}
	}
	return mvc.Response{
		Code:   iris.StatusOK,
		Object: iris.Map{"success": true, "message": "", "result": resultUrl},
	}
}
