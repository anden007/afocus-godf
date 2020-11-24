package controller_manage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/anden007/afocus-godf/src/web/template"
	"github.com/anden007/afocus-godf/src/web/view_model/generator"
	"github.com/anden007/afocus-godf/src/web/view_model/view_model_business"

	"github.com/gobeam/stringy"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/oleiade/reflections"
)

type GenerateController struct {
	Ctx iris.Context
}

func (m *GenerateController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/table/{vueName:string}/{rowNum:int}", "GenerateTable")
	b.Handle("POST", "/tree/{vueName:string}/{rowNum:int}", "GenerateTree")
	b.Handle("GET", "/getEntityData/{typeName:string}", "GetEntityData")
}

func (m *GenerateController) GenerateTable(vueName string, rowNum int) mvc.Result {
	var options generator.Options
	var data iris.Map
	payload, err := m.Ctx.GetBody()
	if err == nil {
		err := json.Unmarshal(payload, &options)
		if err == nil {
			data = iris.Map{
				"component":    Generate("tableIndex", false, vueName, rowNum, options),
				"componentApi": Generate("tableIndex", true, vueName, rowNum, options),
				"add":          Generate("add", false, vueName, rowNum, options),
				"addApi":       Generate("add", true, vueName, rowNum, options),
				"edit":         Generate("edit", false, vueName, rowNum, options),
				"editApi":      Generate("edit", true, vueName, rowNum, options),
				"single":       Generate("table", false, vueName, rowNum, options),
				"singleApi":    Generate("table", true, vueName, rowNum, options),
				"api":          Generate("api", true, vueName, rowNum, options),
			}
		}
	}
	return mvc.Response{
		Object: iris.Map{"success": true, "message": "", "result": data},
	}
}

func (m *GenerateController) GenerateTree(vueName string, rowNum int) mvc.Result {
	var options generator.Options
	var data iris.Map
	payload, err := m.Ctx.GetBody()
	if err == nil {
		err := json.Unmarshal(payload, &options)
		if err == nil {
			data = iris.Map{
				"result":    Generate("tree", false, vueName, rowNum, options),
				"resultApi": Generate("tree", true, vueName, rowNum, options),
				"api":       Generate("treeApi", true, vueName, rowNum, options),
			}
		}
	}
	return mvc.Response{
		Object: iris.Map{"success": true, "message": "", "result": data},
	}
}

func (m *GenerateController) GetEntityData(typeName string) mvc.Result {
	var result []generator.Field
	typ, exists := view_model_business.Types[typeName]
	if exists {
		model := reflect.New(typ).Interface()
		fieldArray, _ := reflections.FieldsDeep(model)
		for index, field := range fieldArray {
			field, _ := reflections.GetFieldTag(model, field, "json")
			fieldCaption, _ := reflections.GetFieldTag(model, field, "caption")
			kind, _ := reflections.GetFieldKind(model, field)
			if field != "id" && field != "-" {
				result = append(result, generator.Field{
					Field:           field,
					Name:            fieldCaption,
					Level:           "2",
					TableShow:       true,
					SortOrder:       float32(index),
					Searchable:      false,
					Editable:        true,
					Type:            GetGenerateType(kind),
					Validate:        false,
					SearchType:      "",
					SearchLevel:     "",
					Sortable:        false,
					DefaultSort:     false,
					DefaultSortType: "",
				})
			}
		}
	} else {
		return mvc.Response{
			Object: iris.Map{"success": false, "message": "", "result": fmt.Sprintf("实体类%s不存在", typeName)},
		}
	}
	jsonString, err := json.MarshalIndent(iris.Map{"data": result}, "", "    ")
	if err == nil {
		return mvc.Response{
			Object: iris.Map{"success": true, "message": "", "result": string(jsonString)},
		}
	} else {
		return mvc.Response{
			Object: iris.Map{"success": false, "message": "", "result": ""},
		}
	}
}

func GetGenerateType(kind reflect.Kind) string {
	result := ""
	switch {
	case kind == reflect.Int64, kind == reflect.Int, kind == reflect.Int8,
		kind == reflect.Int16, kind == reflect.Int32:
		{
			result = "number"
		}
	case kind == reflect.String:
		{
			result = "text"
		}
	case kind == reflect.Bool:
		{
			result = "switch"
		}
	case kind == reflect.Float32, kind == reflect.Float64:
		{
			result = "number"
		}
	case kind == reflect.Struct:
		{
			//time.Time类型会被识别成struct，我要取到这个struct里面的时间值付给dest
			result = "date"
		}
	default:
		result = "text"
	}
	return result
}

func Generate(templateName string, api bool, vueName string, rowNum int, generatorOptions generator.Options) (result string) {
	result = ""
	// 排序

	// 绑定变量
	var firstTwo []generator.Field
	var rest []generator.Field
	vueNameStr := stringy.New(vueName)
	apiName := vueNameStr.CamelCase() //转换成大驼峰

	// 判断有无upload和日期范围搜索等组件
	upload := false
	uploadThumb := false
	wangEditor := false
	quill := false
	password := false
	daterangeSearch := false
	hideSearch := false
	searchSize := 0
	wangEditorWidth := ""
	quillWidth := ""
	toQuill := false
	modalWidth := ""
	width := ""
	editWidth := ""
	itemWidth := ""
	span := 0
	defaultSort := ""
	defaultSortType := ""

	for _, f := range generatorOptions.Fields {
		if "upload" == f.Type {
			upload = true
		}
		if "uploadThumb" == f.Type {
			uploadThumb = true
		}
		if "wangEditor" == f.Type {
			wangEditor = true
		}
		if "quill" == f.Type {
			quill = true
		}
		if "password" == f.Type {
			password = "password" == f.Type
		}
	}

	if "table" == templateName || "tableIndex" == templateName {
		// 判断有无upload和日期范围搜索
		for _, f := range generatorOptions.Fields {
			if f.Searchable && "daterange" == f.SearchType {
				daterangeSearch = true
			}
		}
		// 统计搜索栏个数 判断是否隐藏搜索栏
		count := 0
		for _, f := range generatorOptions.Fields {
			if f.Searchable {
				count++
				if count <= 2 {
					firstTwo = append(firstTwo, f)
				} else {
					rest = append(rest, f)
				}
			}
		}
		if count >= 4 {
			hideSearch = true
		}
		searchSize = count
		// 获取默认排序字段
		for _, f := range generatorOptions.Fields {
			if f.DefaultSort {
				defaultSort = f.Field
				defaultSortType = f.DefaultSortType
				break
			}
		}
	}
	wangEditorWidth = "100%"
	quillWidth = "100%"
	toQuill = false
	// 一行几列
	if rowNum == 1 {
		modalWidth = "500"
		width = "100%"
		editWidth = "100%"
		itemWidth = ""
		span = 9
		if "table" == templateName || "tree" == templateName {
			toQuill = true
		}
		if "add" == templateName || "edit" == templateName {
			width = "570px"
			rowNum = 1
		}
	} else if rowNum == 2 {
		modalWidth = "770"
		width = "250px"
		editWidth = "250px"
		itemWidth = "350px"
		span = 17
		if "table" == templateName || "tree" == templateName {
			toQuill = true
			quillWidth = "610px"
		}
		if "add" == templateName || "edit" == templateName {
			width = "570px"
			rowNum = 1
		}
	} else if rowNum == 3 {
		modalWidth = "980"
		width = "200px"
		editWidth = "200px"
		itemWidth = "300px"
		span = 17
		if "table" == templateName || "tree" == templateName {
			quillWidth = "820px"
			wangEditorWidth = "820px"
		}
		if "add" == templateName || "edit" == templateName {
			width = "570px"
			rowNum = 1
		}
	} else if rowNum == 4 {
		modalWidth = "1130"
		width = "160px"
		editWidth = "160px"
		itemWidth = "260px"
		span = 17
		if "table" == templateName || "tree" == templateName {
			quillWidth = "970px"
			wangEditorWidth = "970px"
		}
		if "add" == templateName || "edit" == templateName {
			width = "570px"
			rowNum = 1
		}
	} else {
		panic("rowNum仅支持数字1-4")
	}
	// 生成代码
	buffer := new(bytes.Buffer)
	switch templateName {
	case "tableIndex":
		options := template.TableIndexOption{
			SearchSize:      searchSize,
			HideSearch:      hideSearch,
			ApiName:         apiName,
			DaterangeSearch: daterangeSearch,
			Api:             false,
			DefaultSort:     defaultSort,
			DefaultSortType: defaultSortType,
		}
		template.TableIndex(generatorOptions.Fields, firstTwo, rest, options, buffer)
		break
	case "edit":
		options := template.EditOption{
			RowNum:          rowNum,
			ItemWidth:       itemWidth,
			Width:           width,
			WangEditorWidth: wangEditorWidth,
			QuillWidth:      quillWidth,
			ApiName:         apiName,
			Api:             api,
			Upload:          upload,
			UploadThumb:     uploadThumb,
			WangEditor:      wangEditor,
			Quill:           quill,
			Password:        password,
		}
		template.Edit(generatorOptions.Fields, options, buffer)
		break
	case "table":
		options := template.TableOption{
			RowNum:          rowNum,
			SearchSize:      searchSize,
			HideSearch:      hideSearch,
			ModalWidth:      modalWidth,
			Width:           width,
			ToQuill:         toQuill,
			WangEditorWidth: wangEditorWidth,
			QuillWidth:      quillWidth,
			ApiName:         apiName,
			Upload:          upload,
			UploadThumb:     uploadThumb,
			WangEditor:      wangEditor,
			Quill:           quill,
			DaterangeSearch: daterangeSearch,
			Password:        password,
			VueName:         vueName,
			Api:             api,
			ItemWidth:       itemWidth,
			DefaultSort:     defaultSort,
			DefaultSortType: defaultSortType,
		}
		template.Table(generatorOptions.Fields, firstTwo, rest, options, buffer)
		break
	case "api":
		template.Api(apiName, generatorOptions.ApiPath, vueName, buffer)
		break
	case "add":
		options := template.AddOption{
			RowNum:          rowNum,
			ItemWidth:       itemWidth,
			Width:           width,
			WangEditorWidth: wangEditorWidth,
			QuillWidth:      quillWidth,
			ApiName:         apiName,
			Api:             api,
			Upload:          upload,
			UploadThumb:     uploadThumb,
			WangEditor:      wangEditor,
			Quill:           quill,
			Password:        password,
		}
		template.Add(generatorOptions.Fields, options, buffer)
		break
	case "tree":
		options := template.TreeOption{
			RowNum:          rowNum,
			SearchSize:      searchSize,
			HideSearch:      hideSearch,
			ModalWidth:      modalWidth,
			Width:           width,
			ToQuill:         toQuill,
			WangEditorWidth: wangEditorWidth,
			QuillWidth:      quillWidth,
			ApiName:         apiName,
			Upload:          upload,
			UploadThumb:     uploadThumb,
			WangEditor:      wangEditor,
			Quill:           quill,
			DaterangeSearch: daterangeSearch,
			Password:        password,
			VueName:         vueName,
			Api:             api,
			ItemWidth:       itemWidth,
			EditWidth:       editWidth,
			Span:            span,
		}
		template.Tree(generatorOptions.Fields, firstTwo, rest, options, buffer)
		break

	case "treeApi":
		template.TreeApi(apiName, generatorOptions.ApiPath, vueName, buffer)
		break
	}

	result = buffer.String()
	return
}
