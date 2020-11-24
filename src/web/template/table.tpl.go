// Code generated by hero.
// source: D:\GoProjects\github.com/anden007/afocus-godf\Backend\src\web\template\table.tpl
// DO NOT EDIT!
package template

import (
	"bytes"
	"github.com/anden007/afocus-godf/src/web/view_model/generator"
)

type TableOption struct {
	RowNum          int
	SearchSize      int
	HideSearch      bool
	ModalWidth      string
	Width           string
	ToQuill         bool
	WangEditorWidth string
	QuillWidth      string
	ApiName         string
	Upload          bool
	UploadThumb     bool
	WangEditor      bool
	Quill           bool
	DaterangeSearch bool
	Password        bool
	VueName         string
	Api             bool
	ItemWidth       string
	DefaultSort     string
	DefaultSortType string
}

func Table(fields []generator.Field, firstTwo []generator.Field, rest []generator.Field, options TableOption, buffer *bytes.Buffer) {
	buffer.WriteString(`
<template>
  <div class="search">
    <Card>
      `)

	if options.SearchSize > 0 && !options.HideSearch {

		buffer.WriteString(`
      <Row `)
		if options.SearchSize > 0 {
			buffer.WriteString(`v-show="openSearch"`)
		}
		buffer.WriteString(` @keydown.enter.native="handleSearch">
        <Form ref="searchForm" :model="searchForm" inline :label-width="70">
        `)

		for _, item := range fields {
			if item.Searchable {

				if item.SearchType == "text" {

					buffer.WriteString(`
            <Form-item label=" `)
					buffer.WriteString(item.Name)
					buffer.WriteString(`" prop="`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`">
              <Input type="text" v-model="searchForm.`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`" placeholder="请输入`)
					buffer.WriteString(item.Name)
					buffer.WriteString(`" clearable style="width: 200px"/>
            </Form-item>
            `)

				}

				if item.SearchType == "select" {

					buffer.WriteString(`
            <Form-item label="`)
					buffer.WriteString(item.Name)
					buffer.WriteString(`" prop="`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`">
              <Select v-model="searchForm.`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`" placeholder="请选择" clearable style="width: 200px">
                <Option value="0">请自行编辑下拉菜单</Option>
              </Select>
            </Form-item>
            `)

				}

				if item.SearchType == "date" {

					buffer.WriteString(`
            <Form-item label="`)
					buffer.WriteString(item.Name)
					buffer.WriteString(`" prop="`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`">
              <DatePicker type="date" v-model="searchForm.`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`" placeholder="请选择" clearable style="width: 200px"></DatePicker>
            </Form-item>
            `)

				}

				if item.SearchType == "daterange" {

					buffer.WriteString(`
            <Form-item label="`)
					buffer.WriteString(item.Name)
					buffer.WriteString(`" prop="`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`">
              <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
            </Form-item>
            `)

				}

				if item.SearchType == "area" {

					buffer.WriteString(`
            <Form-item label="`)
					buffer.WriteString(item.Name)
					buffer.WriteString(`" prop="`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`">
              <al-cascader v-model="searchForm.`)
					buffer.WriteString(item.Field)
					buffer.WriteString(`" data-type="code" level="`)
					buffer.WriteString(item.SearchLevel)
					buffer.WriteString(`" style="width:200px"/>
            </Form-item>
            `)

				}

			}
		}

		buffer.WriteString(`
          <Form-item style="margin-left:-35px;" class="br">
            <Button @click="handleSearch" type="primary" icon="ios-search">搜索</Button>
            <Button @click="handleReset">重置</Button>
          </Form-item>
        </Form>
      </Row>
      `)

	}

	if options.SearchSize > 0 && options.HideSearch {

		buffer.WriteString(`
      <Row @keydown.enter.native="handleSearch">
        <Form ref="searchForm" :model="searchForm" inline :label-width="70" class="search-form">
        `)

		for _, item := range firstTwo {

			if item.SearchType == "text" {

				buffer.WriteString(`
          <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
            <Input type="text" v-model="searchForm.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" placeholder="请输入`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" clearable style="width: 200px"/>
          </Form-item>
          `)

			}

			if item.SearchType == "select" {

				buffer.WriteString(`
          <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
            <Select v-model="searchForm.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" placeholder="请选择" clearable style="width: 200px">
              <Option value="0">请自行编辑下拉菜单</Option>
            </Select>
          </Form-item>
          `)

			}

			if item.SearchType == "date" {

				buffer.WriteString(`
          <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
            <DatePicker type="date" v-model="searchForm.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" placeholder="请选择" clearable style="width: 200px"></DatePicker>
          </Form-item>
          `)

			}

			if item.SearchType == "daterange" {

				buffer.WriteString(`
          <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
            <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
          </Form-item>
          `)

			}

			if item.SearchType == "area" {

				buffer.WriteString(`
          <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
            <al-cascader v-model="searchForm.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" data-type="code" level="`)
				buffer.WriteString(item.SearchLevel)
				buffer.WriteString(`" style="width:200px"/>
          </Form-item>
          `)

			}

		}

		buffer.WriteString(`
          <span v-if="drop">
          `)

		for _, item := range rest {

			if item.SearchType == "text" {

				buffer.WriteString(`
            <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
              <Input type="text" v-model="searchForm.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" placeholder="请输入`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" clearable style="width: 200px"/>
            </Form-item>
            `)

			}

			if item.SearchType == "select" {

				buffer.WriteString(`
            <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
              <Select v-model="searchForm.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" placeholder="请选择" clearable style="width: 200px">
                <Option value="0">请自行编辑下拉菜单</Option>
              </Select>
            </Form-item>
            `)

			}

			if item.SearchType == "date" {

				buffer.WriteString(`
            <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
              <DatePicker type="date" v-model="searchForm.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" placeholder="请选择" clearable style="width: 200px"></DatePicker>
            </Form-item>
            `)

			}

			if item.SearchType == "daterange" {

				buffer.WriteString(`
            <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
              <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
            </Form-item>
            `)

			}

			if item.SearchType == "area" {

				buffer.WriteString(`
            <Form-item label="`)
				buffer.WriteString(item.Name)
				buffer.WriteString(`" prop="`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
              <al-cascader v-model="searchForm.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" data-type="code" level="`)
				buffer.WriteString(item.SearchLevel)
				buffer.WriteString(`" style="width:200px"/>
            </Form-item>
            `)

			}

		}

		buffer.WriteString(`
          </span>
          <Form-item style="margin-left:-35px;" class="br">
            <Button @click="handleSearch" type="primary" icon="ios-search">搜索</Button>
            <Button @click="handleReset">重置</Button>
            <a class="drop-down" @click="dropDown">
              {{dropDownContent}}
              <Icon :type="dropDownIcon"></Icon>
            </a>
          </Form-item>
        </Form>
      </Row>
      `)

	}

	buffer.WriteString(`
      <Row class="operation">
        <Button @click="add" type="primary" icon="md-add">添加</Button>
        <Button @click="delAll" icon="md-trash">批量删除</Button>
        <Button @click="getDataList" icon="md-refresh">刷新</Button>
        `)
	if options.SearchSize > 0 {
		buffer.WriteString(`
        <Button type="dashed" @click="openSearch=!openSearch">{{openSearch ? "关闭搜索" : "开启搜索"}}</Button>
        `)
	}
	buffer.WriteString(`
        <Button type="dashed" @click="openTip=!openTip">{{openTip ? "关闭提示" : "开启提示"}}</Button>
      </Row>
      <Row v-show="openTip">
        <Alert show-icon>
          已选择 <span class="select-count">{{selectCount}}</span> 项
          <a class="select-clear" @click="clearSelectAll">清空</a>
        </Alert>
      </Row>
      <Row>
        <Table :loading="loading" border :columns="columns" :data="data" ref="table" sortable="custom" @on-sort-change="changeSort" @on-selection-change="changeSelect"></Table>
      </Row>
      <Row type="flex" justify="end" class="page">
        <Page :current="searchForm.pageNumber" :total="total" :page-size="searchForm.pageSize" @on-change="changePage" @on-page-size-change="changePageSize" :page-size-opts="[10,20,50]" size="small" show-total show-elevator show-sizer></Page>
      </Row>
    </Card>
    <Modal :title="modalTitle" v-model="modalVisible" :mask-closable='false' :width="`)
	buffer.WriteString(options.ModalWidth)
	buffer.WriteString(`">
      <Form ref="form" :model="form" :label-width="100" :rules="formValidate" label-position="left" `)
	if options.RowNum > 1 {
		buffer.WriteString(`inline`)
	}
	buffer.WriteString(`>
        `)

	for _, item := range fields {
		if item.Editable {

			buffer.WriteString(`
        <FormItem label="`)
			buffer.WriteString(item.Name)
			buffer.WriteString(`" prop="`)
			buffer.WriteString(item.Field)
			buffer.WriteString(`" `)
			if options.RowNum > 1 && (item.Type == "switch" || item.Type == "radio") {
				buffer.WriteString(`style="width:`)
				buffer.WriteString(options.ItemWidth)
				buffer.WriteString(`"`)
			}
			if item.Type == "wangEditor" || item.Type == "quill" {
				buffer.WriteString(`class="form-`)
				buffer.WriteString(item.Type)
				buffer.WriteString(`"`)
			}
			buffer.WriteString(`>
          `)

			if item.Type == "text" {

				buffer.WriteString(`
          <Input v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" clearable style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"/>
          `)

			}

			if item.Type == "select" {

				buffer.WriteString(`
          <Select v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" clearable style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`">
            <Option value="0">请自行编辑下拉菜单</Option>
          </Select>
          `)

			}

			if item.Type == "switch" {

				buffer.WriteString(`
          <i-switch v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`"></i-switch>
          `)

			}

			if item.Type == "radio" {

				buffer.WriteString(`
          <RadioGroup v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`">
            <Radio label="0">请自行编辑单选框</Radio>
            <Radio label="1">请自行编辑单选框</Radio>
          </RadioGroup>
          `)

			}

			if item.Type == "number" {

				buffer.WriteString(`
          <InputNumber v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"></InputNumber>
          `)

			}

			if item.Type == "date" {

				buffer.WriteString(`
          <DatePicker type="date" v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" clearable style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"></DatePicker>
          `)

			}

			if item.Type == "daterange" {

				buffer.WriteString(`
          <DatePicker type="daterange" v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" clearable style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"></DatePicker>
          `)

			}

			if item.Type == "time" {

				buffer.WriteString(`
          <TimePicker type="time" v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" clearable style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"></TimePicker>
          `)

			}

			if item.Type == "area" {

				buffer.WriteString(`
          <al-cascader v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" data-type="code" level="`)
				buffer.WriteString(item.Level)
				buffer.WriteString(`" style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"/>
          `)

			}

			if item.Type == "slider" {

				buffer.WriteString(`
          <Slider v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"></Slider>
          `)

			}

			if item.Type == "upload" {

				buffer.WriteString(`
          <upload-pic-input v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"></upload-pic-input>
          `)

			}

			if item.Type == "uploadThumb" {

				buffer.WriteString(`
          <uploadThumb v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" multiple style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"></uploadThumb>
          `)

			}

			if item.Type == "wangEditor" && !options.ToQuill {

				buffer.WriteString(`
          <wangEditor v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" style="width:`)
				buffer.WriteString(options.WangEditorWidth)
				buffer.WriteString(`"></wangEditor>
          `)

			}

			if item.Type == "quill" || (item.Type == "wangEditor" && options.ToQuill) {

				buffer.WriteString(`
          <quill v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" style="width:`)
				buffer.WriteString(options.QuillWidth)
				buffer.WriteString(`"></quill>
          `)

			}

			if item.Type == "password" {

				buffer.WriteString(`
          <password v-model="form.`)
				buffer.WriteString(item.Field)
				buffer.WriteString(`" style="width:`)
				buffer.WriteString(options.Width)
				buffer.WriteString(`"></password>
          `)

			}

			buffer.WriteString(`
        </FormItem>
        `)

		}
	}

	buffer.WriteString(`
      </Form>
      <div slot="footer">
        <Button type="text" @click="modalVisible=false">取消</Button>
        <Button type="primary" :loading="submitLoading" @click="handleSubmit">提交</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
`)

	if options.Api {

		buffer.WriteString(`
// 根据你的实际请求api.js位置路径修改
import { get`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(`List, add`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(`, edit`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(`, delete`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(` } from "@/api/index";
import moment from "moment";
`)

	}

	if options.Upload {

		buffer.WriteString(`
import uploadPicInput from "@/views/my-components/xboot/upload-pic-input";
`)

	}

	if options.UploadThumb {

		buffer.WriteString(`
import uploadThumb from "@/views/my-components/xboot/upload-pic-thumb";
`)

	}

	if options.WangEditor && !options.ToQuill {

		buffer.WriteString(`
import wangEditor from "@/views/my-components/xboot/editor";
`)

	}

	if options.Quill || options.ToQuill {

		buffer.WriteString(`
import quill from "@/views/my-components/xboot/quill";
`)

	}

	if options.Password {

		buffer.WriteString(`
import password from "@/views/my-components/xboot/set-password";
`)

	}

	buffer.WriteString(`
export default {
  name: "`)
	buffer.WriteString(options.VueName)
	buffer.WriteString(`",
  components: {
    `)

	if options.Upload {

		buffer.WriteString(`
    uploadPicInput,
    `)

	}

	if options.UploadThumb {

		buffer.WriteString(`
    uploadThumb,
    `)

	}

	if options.WangEditor && !options.ToQuill {

		buffer.WriteString(`
    wangEditor,
    `)

	}

	if options.Quill || options.ToQuill {

		buffer.WriteString(`
    quill,
    `)

	}

	if options.Password {

		buffer.WriteString(`
    password,
    `)

	}

	buffer.WriteString(`
  },
  data() {
    return {
      `)
	if options.SearchSize > 0 {
		buffer.WriteString(`
      openSearch: true, // 显示搜索
      `)
	}
	buffer.WriteString(`
      openTip: true, // 显示提示
      loading: true, // 表单加载状态
      modalType: 0, // 添加或编辑标识
      modalVisible: false, // 添加或编辑显示
      modalTitle: "", // 添加或编辑标题
      `)
	if options.HideSearch {
		buffer.WriteString(`
      drop: false,
      dropDownContent: "展开",
      dropDownIcon: "ios-arrow-down",
      `)
	}
	buffer.WriteString(`
      searchForm: { // 搜索框初始化对象
        pageNumber: 1, // 当前页数
        pageSize: 10, // 页面大小
        sort: "`)
	buffer.WriteString(options.DefaultSort)
	buffer.WriteString(`", // 默认排序字段
        order: "`)
	buffer.WriteString(options.DefaultSortType)
	buffer.WriteString(`", // 默认排序方式
        `)
	if options.DaterangeSearch {
		buffer.WriteString(`
        startDate: "", // 起始时间
        endDate: "" // 终止时间
        `)
	}
	buffer.WriteString(`
      },
      `)
	if options.DaterangeSearch {
		buffer.WriteString(`
      selectDate: null,
      `)
	}
	buffer.WriteString(`
      form: { // 添加或编辑表单对象初始化数据
        `)

	for _, item := range fields {
		if item.Editable {

			if item.Type == "switch" {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: true,
        `)
			} else if item.Type == "number" || item.Type == "slider" {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: 0,
        `)
			} else if item.Type == "area" || item.Type == "uploadThumb" {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: [],
        `)
			} else {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: "",
        `)
			}

		}
	}

	buffer.WriteString(`
      },
      // 表单验证规则
      formValidate: {
        `)

	for _, item := range fields {
		if item.Editable && item.Validate {

			if item.Type == "daterange" || item.Type == "area" || item.Type == "uploadThumb" {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: [{ type: "array", required: true, message: "不能为空", trigger: "blur" }],
        `)
			} else if item.Type == "date" {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: [{ required: true, message: "不能为空", trigger: "blur", pattern: /.+/ }],
        `)
			} else if item.Type == "number" || item.Type == "slider" {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: [{ type: "number", required: true, message: "不能为空", trigger: "blur" }],
        `)
			} else if item.Type == "switch" {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: [{ type: "boolean", required: true, message: "不能为空", trigger: "blur" }],
        `)
			} else {
				buffer.WriteString(item.Field)
				buffer.WriteString(`: [{ required: true, message: "不能为空", trigger: "blur" }],
        `)
			}

		}
	}

	buffer.WriteString(`
      },
      submitLoading: false, // 添加或编辑提交状态
      selectList: [], // 多选数据
      selectCount: 0, // 多选计数
      columns: [
        // 表头
        {
          type: "selection",
          width: 60,
          align: "center"
        },
        {
          type: "index",
          width: 60,
          align: "center"
        },
        `)

	for _, item := range fields {
		if item.TableShow {

			buffer.WriteString(`
        {
          title: "`)
			buffer.WriteString(item.Name)
			buffer.WriteString(`",
          key: "`)
			buffer.WriteString(item.Field)
			buffer.WriteString(`",
          minWidth: 120,
          `)

			if item.Sortable {

				buffer.WriteString(`
          sortable: true,
          `)

			} else {

				buffer.WriteString(`
          sortable: false,
          `)

			}

			if item.DefaultSort {

				buffer.WriteString(`
          sortType: "`)
				buffer.WriteString(item.DefaultSortType)
				buffer.WriteString(`"
          `)

			}

			buffer.WriteString(`
        },
        `)

		}
	}

	buffer.WriteString(`
        {
          title: "操作",
          key: "action",
          align: "center",
          width: 200,
          render: (h, params) => {
            return h("div", [
              h(
                "Button",
                {
                  props: {
                    type: "primary",
                    size: "small",
                    icon: "ios-create-outline"
                  },
                  style: {
                    marginRight: "5px"
                  },
                  on: {
                    click: () => {
                      this.edit(params.row);
                    }
                  }
                },
                "编辑"
              ),
              h(
                "Button",
                {
                  props: {
                    type: "error",
                    size: "small",
                    icon: "md-trash"
                  },
                  on: {
                    click: () => {
                      this.remove(params.row);
                    }
                  }
                },
                "删除"
              )
            ]);
          }
        }
      ],
      data: [], // 表单数据
      total: 0 // 表单数据总数
    };
  },
  methods: {
    init() {
      this.getDataList();
    },
    changePage(v) {
      this.searchForm.pageNumber = v;
      this.getDataList();
      this.clearSelectAll();
    },
    changePageSize(v) {
      this.searchForm.pageSize = v;
      this.getDataList();
    },
    `)
	if options.SearchSize > 0 {
		buffer.WriteString(`
    handleSearch() {
      this.searchForm.pageNumber = 1;
      this.searchForm.pageSize = 10;
      this.getDataList();
    },
    `)
	}
	buffer.WriteString(`
    handleReset() {
      this.$refs.searchForm.resetFields();
      this.searchForm.pageNumber = 1;
      this.searchForm.pageSize = 10;
      `)
	if options.DaterangeSearch {
		buffer.WriteString(`
      this.selectDate = null;
      this.searchForm.startDate = "";
      this.searchForm.endDate = "";
      `)
	}
	buffer.WriteString(`
      // 重新加载数据
      this.getDataList();
    },
    changeSort(e) {
      this.searchForm.sort = e.key;
      this.searchForm.order = e.order;
      if (e.order === "normal") {
        this.searchForm.order = "";
      }
      this.getDataList();
    },
    clearSelectAll() {
      this.$refs.table.selectAll(false);
    },
    changeSelect(e) {
      this.selectList = e;
      this.selectCount = e.length;
    },
    `)
	if options.DaterangeSearch {
		buffer.WriteString(`
    selectDateRange(v) {
      if (v) {
        this.searchForm.startDate = moment(v[0]).format("YYYY-MM-DD HH:mm:ss");
        this.searchForm.endDate = moment(v[1]).format("YYYY-MM-DD HH:mm:ss");
      }
    },
    `)
	}
	if options.HideSearch {
		buffer.WriteString(`
    dropDown() {
      if (this.drop) {
        this.dropDownContent = "展开";
        this.dropDownIcon = "ios-arrow-down";
      } else {
        this.dropDownContent = "收起";
        this.dropDownIcon = "ios-arrow-up";
      }
      this.drop = !this.drop;
    },
    `)
	}
	buffer.WriteString(`
    getDataList() {
      this.loading = true;
      `)

	if options.Api {

		buffer.WriteString(`
      get`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(`List(this.searchForm).then(res => {
        this.loading = false;
        if (res.success) {
          if(!res.result.content){
            this.data =[];
          }else{
            this.data = res.result.content;
          }
          this.total = res.result.totalElements;
        }
      });
      `)

	} else {

		buffer.WriteString(`
      // 带多条件搜索参数获取表单数据 请自行修改接口
      // this.getRequest("请求路径", this.searchForm).then(res => {
      //   this.loading = false;
      //   if (res.success) {
      //     this.data = res.result.content;
      //     this.total = res.result.totalElements;
      //   }
      // });
      // 以下为模拟数据
      //this.data = [
      //];
      this.total = this.data.length;
      this.loading = false;
      `)

	}

	buffer.WriteString(`
    },
    handleSubmit() {
      this.$refs.form.validate(valid => {
        if (valid) {
        `)
	for _, item := range fields {
		if item.Type == "date" {
			buffer.WriteString(`
          this.form.`)
			buffer.WriteString(item.Field)
			buffer.WriteString(` = moment(this.form.`)
			buffer.WriteString(item.Field)
			buffer.WriteString(`).format("YYYY-MM-DD HH:mm:ss");
        `)
		}
	}
	buffer.WriteString(`
          this.submitLoading = true;
          if (this.modalType === 0) {
            // 添加 避免编辑后传入id等数据 记得删除
            delete this.form.id;
            `)

	if options.Api {
		buffer.WriteString(`
            add`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(`(this.form).then(res => {
              this.submitLoading = false;
              if (res.success) {
                this.$Message.success("操作成功");
                this.getDataList();
                this.modalVisible = false;
              }
            });
            `)

	} else {

		buffer.WriteString(`
            // this.postRequest("请求地址", this.form).then(res => {
            //   this.submitLoading = false;
            //   if (res.success) {
            //     this.$Message.success("操作成功");
            //     this.getDataList();
            //     this.modalVisible = false;
            //   }
            // });
            // 模拟请求成功
            this.submitLoading = false;
            this.$Message.success("操作成功");
            this.getDataList();
            this.modalVisible = false;
            `)

	}

	buffer.WriteString(`
          } else {
            // 编辑
            `)

	if options.Api {

		buffer.WriteString(`
            edit`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(`(this.form).then(res => {
              this.submitLoading = false;
              if (res.success) {
                this.$Message.success("操作成功");
                this.getDataList();
                this.modalVisible = false;
              }
            });
            `)

	} else {

		buffer.WriteString(`
            // this.postRequest("请求地址", this.form).then(res => {
            //   this.submitLoading = false;
            //   if (res.success) {
            //     this.$Message.success("操作成功");
            //     this.getDataList();
            //     this.modalVisible = false;
            //   }
            // });
            // 模拟请求成功
            this.submitLoading = false;
            this.$Message.success("操作成功");
            this.getDataList();
            this.modalVisible = false;
            `)

	}

	buffer.WriteString(`
          }
        }
      });
    },
    add() {
      this.modalType = 0;
      this.modalTitle = "添加";
      this.$refs.form.resetFields();
      delete this.form.id;
      this.modalVisible = true;
    },
    edit(v) {
      this.modalType = 1;
      this.modalTitle = "编辑";
      this.$refs.form.resetFields();
      // 转换null为""
      for (let attr in v) {
        if (v[attr] === null) {
          v[attr] = "";
        }
      }
      let str = JSON.stringify(v);
      let data = JSON.parse(str);
      this.form = data;
      this.modalVisible = true;
    },
    remove(v) {
      this.$Modal.confirm({
        title: "确认删除",
        // 记得确认修改此处
        content: "您确认要删除 " + v.name + " ?",
        loading: true,
        onOk: () => {
          // 删除
          `)

	if options.Api {

		buffer.WriteString(`
          delete`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(`(v.id).then(res => {
            this.$Modal.remove();
            if (res.success) {
              this.$Message.success("操作成功");
              this.getDataList();
            }
          });
          `)

	} else {

		buffer.WriteString(`
          // this.deleteRequest("请求地址，如/deleteByIds/" + v.id).then(res => {
          //   this.$Modal.remove();
          //   if (res.success) {
          //     this.$Message.success("操作成功");
          //     this.getDataList();
          //   }
          // });
          // 模拟请求成功
          this.$Message.success("操作成功");
          this.$Modal.remove();
          this.getDataList();
          `)

	}

	buffer.WriteString(`
        }
      });
    },
    delAll() {
      if (this.selectCount <= 0) {
        this.$Message.warning("您还未选择要删除的数据");
        return;
      }
      this.$Modal.confirm({
        title: "确认删除",
        content: "您确认要删除所选的 " + this.selectCount + " 条数据?",
        loading: true,
        onOk: () => {
          let ids = "";
          this.selectList.forEach(function(e) {
            ids += e.id + ",";
          });
          ids = ids.substring(0, ids.length - 1);
          // 批量删除
          `)

	if options.Api {

		buffer.WriteString(`
          delete`)
		buffer.WriteString(options.ApiName)
		buffer.WriteString(`(ids).then(res => {
            this.$Modal.remove();
            if (res.success) {
              this.$Message.success("操作成功");
              this.clearSelectAll();
              this.getDataList();
            }
          });
          `)

	} else {

		buffer.WriteString(`
          // this.deleteRequest("请求地址，如/deleteByIds/" + ids).then(res => {
          //   this.$Modal.remove();
          //   if (res.success) {
          //     this.$Message.success("操作成功");
          //     this.clearSelectAll();
          //     this.getDataList();
          //   }
          // });
          // 模拟请求成功
          this.$Message.success("操作成功");
          this.$Modal.remove();
          this.clearSelectAll();
          this.getDataList();
          `)

	}

	buffer.WriteString(`
        }
      });
    }
  },
  mounted() {
    this.init();
  }
};
</script>
<style lang="less">
// 建议引入通用样式 具体路径自行修改 可删除下面样式代码
// @import "../../../styles/table-common.less";
.search {
    .operation {
        margin-bottom: 2vh;
    }
    .select-count {
        font-weight: 600;
        color: #40a9ff;
    }
    .select-clear {
        margin-left: 10px;
    }
    .page {
        margin-top: 2vh;
    }
    .drop-down {
        margin-left: 5px;
    }
}
</style>`)

}