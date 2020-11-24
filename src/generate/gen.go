package generate

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/anden007/afocus-godf/src/generate/template"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_business"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/gobeam/stringy"
)

func GenCode(pkgCode string, model string, target string) {
	if pkgCode != "" && model != "" {
		if target == "all" {
			str := stringy.New(model)
			fileName := str.SnakeCase().ToLower()
			hasCreateTime, _ := checkCreateTimeField(pkgCode, model)
			//生成接口
			buffer := new(bytes.Buffer)
			template.InterFace(pkgCode, model, buffer)
			WriteWithFileWrite(fmt.Sprintf("./src/interfaces/interface_%s/", pkgCode), fmt.Sprintf("%s.go", fileName), buffer.String())
			//生成仓库
			buffer = new(bytes.Buffer)
			template.Repository(pkgCode, model, hasCreateTime, buffer)
			WriteWithFileWrite(fmt.Sprintf("./src/repositories/repository_%s/", pkgCode), fmt.Sprintf("%s.go", fileName), buffer.String())
			//生成服务
			buffer = new(bytes.Buffer)
			template.Service(pkgCode, model, buffer)
			WriteWithFileWrite(fmt.Sprintf("./src/services/service_%s/", pkgCode), fmt.Sprintf("%s.go", fileName), buffer.String())
			//生成控制器
			buffer = new(bytes.Buffer)
			template.Controller(pkgCode, model, buffer)
			WriteWithFileWrite(fmt.Sprintf("./src/web/controllers/controller_%s/", pkgCode), fmt.Sprintf("%s.go", fileName), buffer.String())
			//生成视图模型
			buffer = new(bytes.Buffer)
			template.ViewModel(pkgCode, model, buffer)
			WriteWithFileWrite(fmt.Sprintf("./src/web/view_model/view_model_%s/", pkgCode), fmt.Sprintf("%s.go", fileName), buffer.String())
			//生成模型静态字段
			var typ reflect.Type = nil
			var exists bool = false
			if pkgCode == "business" {
				typ, exists = model_business.Types[model]
			} else if pkgCode == "manage" {
				typ, exists = model_manage.Types[model]
			}
			if exists {
				modelType := reflect.New(typ).Interface()
				if dbFields, fieldErr := lib.GetModelDBFieldNames(modelType); fieldErr == nil {
					buffer = new(bytes.Buffer)
					template.Model(pkgCode, model, dbFields, buffer)
					WriteWithFileWrite(fmt.Sprintf("./src/model/model_%s/", pkgCode), fmt.Sprintf("%s_fields.go", fileName), buffer.String())
				}
			} else {
				fmt.Printf("× 生成模型%s静态字段失败, pkgreflect.go文件中不存在相关字段！\r\n", model)
			}
			//生成DI注入器
			buffer = new(bytes.Buffer)
			template.DI(pkgCode, model, buffer)
			WriteWithFileWrite("./src/generate/di/", fmt.Sprintf("%s-%s.go", pkgCode, fileName), buffer.String())

			replaceService(pkgCode, model)
			replaceDB(pkgCode, model)
		} else if target == "fields" {
			//生成模型静态字段
			str := stringy.New(model)
			fileName := str.SnakeCase().ToLower()
			var typ reflect.Type = nil
			var exists bool = false
			if pkgCode == "business" {
				typ, exists = model_business.Types[model]
			} else if pkgCode == "manage" {
				typ, exists = model_manage.Types[model]
			}
			if exists {
				modelType := reflect.New(typ).Interface()
				if dbFields, fieldErr := lib.GetModelDBFieldNames(modelType); fieldErr == nil {
					buffer := new(bytes.Buffer)
					template.Model(pkgCode, model, dbFields, buffer)
					WriteWithFileWrite(fmt.Sprintf("./src/model/model_%s/", pkgCode), fmt.Sprintf("%s_fields.go", fileName), buffer.String())
				}
			} else {
				fmt.Printf("× 生成模型%s静态字段失败, pkgreflect.go文件中不存在相关字段！\r\n", model)
			}
		}
	} else {
		panic("生成失败，pkgCode或model参数无效")
	}
}

func WriteWithFileWrite(filePath, fileName, content string) {
	targetFile := filePath + fileName
	if exist, _ := PathExists(filePath); !exist {
		err := os.Mkdir(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("× 生成目录失败，请检查文件路径", err.Error())
			return
		}
	}
	fileObj, err := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("× 生成文件失败，请检查文件路径", err.Error())
	}
	defer fileObj.Close()
	if fileObj != nil {
		if _, err := fileObj.WriteString(content); err == nil {
			fmt.Printf("√ %s  --生成成功\n", targetFile)
			return
		}
	}
	fmt.Printf("× %s  --生成失败\n", targetFile)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func replaceService(pkgCode, model string) {
	serviceFile := "./src/interfaces/service_center.go"
	importCode := fmt.Sprintf("interface_%s", pkgCode)
	serviceFun := fmt.Sprintf("Get%sService", model)
	funExists := false
	importExists := false
	f, err := os.Open(serviceFile)
	if f != nil && err == nil {
		defer f.Close()
		buf := bufio.NewReader(f)
		var result = ""
		for {
			hasReplace := false
			line, _, err := buf.ReadLine()
			if err == io.EOF {
				break
			}
			if !importExists && strings.Contains(string(line), importCode) {
				importExists = true
			}
			if !importExists && strings.Contains(string(line), "<AUTO_GENERATE_IMPORT_CODE_HERE>") {
				hasReplace = true
				result += fmt.Sprintf("\t\"github.com/anden007/afocus-godf/src/interfaces/interface_%s\"\n\t// <AUTO_GENERATE_IMPORT_CODE_HERE>自动生成代码必须的标记，请勿删除！！！\n", pkgCode)
			}
			if !funExists && strings.Contains(string(line), serviceFun) {
				funExists = true
			}
			if !funExists && strings.Contains(string(line), "<AUTO_GENERATE_SERVICE_CODE_HERE>") {
				hasReplace = true
				result += fmt.Sprintf("\tGet%sService func() interface_%s.I%sService\n\t// <AUTO_GENERATE_SERVICE_CODE_HERE>自动生成代码必须的标记，请勿删除！！！\n", model, pkgCode, model)
			}
			if !hasReplace {
				result += string(line) + "\n"
			}
		}
		fw, err := os.OpenFile(serviceFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
		if err == nil {
			w := bufio.NewWriter(fw)
			_, err = w.WriteString(result)
			if err != nil {
				panic(err)
			}
			w.Flush()
		}
	}
}

func replaceDB(pkgCode, model string) {
	targetFile := "./src/generate/model.go"
	importCode := fmt.Sprintf("model_%s", pkgCode)
	generateCode := fmt.Sprintf("&model_%s.%s{},", pkgCode, model)
	codeExists := false
	importExists := false
	f, err := os.Open(targetFile)
	if f != nil && err == nil {
		defer f.Close()
		buf := bufio.NewReader(f)
		var result = ""
		for {
			hasReplace := false
			line, _, err := buf.ReadLine()
			if err == io.EOF {
				break
			}
			if !importExists && strings.Contains(string(line), importCode) {
				importExists = true
			}
			if !importExists && strings.Contains(string(line), "<AUTO_GENERATE_IMPORT_CODE_HERE>") {
				hasReplace = true
				result += fmt.Sprintf("\t\"github.com/anden007/afocus-godf/src/model/model_%s\"\n\t// <AUTO_GENERATE_IMPORT_CODE_HERE>自动生成代码必须的标记，请勿删除！！！\n", pkgCode)
			}
			if !codeExists && strings.Contains(string(line), generateCode) {
				codeExists = true
			}
			if !codeExists && strings.Contains(string(line), "<AUTO_GENERATE_CODE_HERE>") {
				hasReplace = true
				result += fmt.Sprintf("\t\t\t&model_%s.%s{},\n\t\t\t// <AUTO_GENERATE_CODE_HERE>自动生成代码必须的标记，请勿删除！！！\n", pkgCode, model)
			}
			if !hasReplace {
				result += string(line) + "\n"
			}
		}
		fw, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
		if err == nil {
			w := bufio.NewWriter(fw)
			_, err = w.WriteString(result)
			if err != nil {
				panic(err)
			}
			w.Flush()
		}
	}
}

func checkCreateTimeField(pkgCode, model string) (result bool, err error) {
	result = false
	str := stringy.New(model)
	fileName := str.SnakeCase().ToLower()
	targetFile := fmt.Sprintf("./src/model/model_%s/%s.go", pkgCode, fileName)
	f, err := os.Open(targetFile)
	if f != nil && err == nil {
		defer f.Close()
		if err != nil {
			return
		}
		buf := bufio.NewReader(f)
		for {
			line, _, err := buf.ReadLine()
			if err == io.EOF {
				break
			}
			if strings.Contains(string(line), "CreateTime") {
				result = true
				break
			}
		}
	}
	return
}
