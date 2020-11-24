<%!
import (
  "github.com/anden007/afocus-godf/src/web/view_model/generator"
)

type TableOption struct{
  RowNum int
  SearchSize int
  HideSearch bool
  ModalWidth string
  Width string
  ToQuill bool
  WangEditorWidth string
  QuillWidth string
  ApiName string
  Upload bool
  UploadThumb bool
  WangEditor bool
  Quill bool
  DaterangeSearch bool
  Password bool
  VueName string
  Api bool
  ItemWidth string
  DefaultSort string
  DefaultSortType string
}
%>
<%: func Table(fields []generator.Field, firstTwo []generator.Field, rest []generator.Field,options TableOption, buffer *bytes.Buffer) %>
<template>
  <div class="search">
    <Card>
      <%
      if options.SearchSize>0&&!options.HideSearch{
      %>
      <Row <% if options.SearchSize>0{ %>v-show="openSearch"<% } %> @keydown.enter.native="handleSearch">
        <Form ref="searchForm" :model="searchForm" inline :label-width="70">
        <%
        for _, item := range fields {
          if item.Searchable {
        %>
            <%
            if item.SearchType=="text" {
            %>
            <Form-item label=" <%==s item.Name%>" prop="<%==s item.Field%>">
              <Input type="text" v-model="searchForm.<%==s item.Field%>" placeholder="请输入<%==s item.Name%>" clearable style="width: 200px"/>
            </Form-item>
            <%
            }
            %>
            <%
            if item.SearchType=="select"{
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <Select v-model="searchForm.<%==s item.Field%>" placeholder="请选择" clearable style="width: 200px">
                <Option value="0">请自行编辑下拉菜单</Option>
              </Select>
            </Form-item>
            <%
            }
            %>
            <%
            if item.SearchType=="date" {
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <DatePicker type="date" v-model="searchForm.<%==s item.Field%>" placeholder="请选择" clearable style="width: 200px"></DatePicker>
            </Form-item>
            <%
            }
            %>
            <%
            if item.SearchType=="daterange" {
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
            </Form-item>
            <%
            }
            %>
            <%
            if item.SearchType=="area"{
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <al-cascader v-model="searchForm.<%==s item.Field%>" data-type="code" level="<%==s item.SearchLevel%>" style="width:200px"/>
            </Form-item>
            <%
            }
            %>
        <%
          }
        }
        %>
          <Form-item style="margin-left:-35px;" class="br">
            <Button @click="handleSearch" type="primary" icon="ios-search">搜索</Button>
            <Button @click="handleReset">重置</Button>
          </Form-item>
        </Form>
      </Row>
      <%
      }
      %>
      <%
      if options.SearchSize > 0 && options.HideSearch{
      %>
      <Row @keydown.enter.native="handleSearch">
        <Form ref="searchForm" :model="searchForm" inline :label-width="70" class="search-form">
        <%
        for _, item := range firstTwo {
        %>
          <%
          if item.SearchType=="text" {
          %>
          <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
            <Input type="text" v-model="searchForm.<%==s item.Field%>" placeholder="请输入<%==s item.Name%>" clearable style="width: 200px"/>
          </Form-item>
          <%
          }
          %>
          <%
          if item.SearchType=="select" {
          %>
          <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
            <Select v-model="searchForm.<%==s item.Field%>" placeholder="请选择" clearable style="width: 200px">
              <Option value="0">请自行编辑下拉菜单</Option>
            </Select>
          </Form-item>
          <%
          }
          %>
          <%
          if item.SearchType=="date" {
          %>
          <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
            <DatePicker type="date" v-model="searchForm.<%==s item.Field%>" placeholder="请选择" clearable style="width: 200px"></DatePicker>
          </Form-item>
          <%
          }
          %>
          <%
          if item.SearchType=="daterange" {
          %>
          <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
            <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
          </Form-item>
          <%
          }
          %>
          <%
          if item.SearchType=="area" {
          %>
          <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
            <al-cascader v-model="searchForm.<%==s item.Field%>" data-type="code" level="<%==s item.SearchLevel%>" style="width:200px"/>
          </Form-item>
          <%
          }
          %>
        <%
        }
        %>
          <span v-if="drop">
          <%
          for _, item := range rest {
          %>
            <%
            if item.SearchType=="text" {
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <Input type="text" v-model="searchForm.<%==s item.Field%>" placeholder="请输入<%==s item.Name%>" clearable style="width: 200px"/>
            </Form-item>
            <%
            }
            %>
            <%
            if item.SearchType=="select" {
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <Select v-model="searchForm.<%==s item.Field%>" placeholder="请选择" clearable style="width: 200px">
                <Option value="0">请自行编辑下拉菜单</Option>
              </Select>
            </Form-item>
            <%
            }
            %>
            <%
            if item.SearchType=="date" {
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <DatePicker type="date" v-model="searchForm.<%==s item.Field%>" placeholder="请选择" clearable style="width: 200px"></DatePicker>
            </Form-item>
            <%
            }
            %>
            <%
            if item.SearchType=="daterange" {
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
            </Form-item>
            <%
            }
            %>
            <%
            if item.SearchType=="area" {
            %>
            <Form-item label="<%==s item.Name%>" prop="<%==s item.Field%>">
              <al-cascader v-model="searchForm.<%==s item.Field%>" data-type="code" level="<%==s item.SearchLevel%>" style="width:200px"/>
            </Form-item>
            <%
            }
            %>
          <%
          }
          %>
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
      <%
      }
      %>
      <Row class="operation">
        <Button @click="add" type="primary" icon="md-add">添加</Button>
        <Button @click="delAll" icon="md-trash">批量删除</Button>
        <Button @click="getDataList" icon="md-refresh">刷新</Button>
        <% if options.SearchSize>0 { %>
        <Button type="dashed" @click="openSearch=!openSearch">{{openSearch ? "关闭搜索" : "开启搜索"}}</Button>
        <% } %>
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
    <Modal :title="modalTitle" v-model="modalVisible" :mask-closable='false' :width="<%==s options.ModalWidth%>">
      <Form ref="form" :model="form" :label-width="100" :rules="formValidate" label-position="left" <% if options.RowNum > 1{ %>inline<% } %>>
        <%
        for _, item := range fields {
          if item.Editable {
        %>
        <FormItem label="<%==s item.Name%>" prop="<%==s item.Field%>" <% if options.RowNum > 1 && (item.Type == "switch"||item.Type == "radio"){ %>style="width:<%==s options.ItemWidth%>"<% } %> <% if item.Type=="wangEditor"||item.Type=="quill"{ %>class="form-<%==s  item.Type%>"<% } %>>
          <%
          if item.Type=="text" {
          %>
          <Input v-model="form.<%==s item.Field%>" clearable style="width:<%==s options.Width%>"/>
          <%
          }
          %>
          <%
          if item.Type=="select" {
          %>
          <Select v-model="form.<%==s item.Field%>" clearable style="width:<%==s options.Width%>">
            <Option value="0">请自行编辑下拉菜单</Option>
          </Select>
          <%
          }
          %>
          <%
          if item.Type=="switch" {
          %>
          <i-switch v-model="form.<%==s item.Field%>"></i-switch>
          <%
          }
          %>
          <%
          if item.Type=="radio" {
          %>
          <RadioGroup v-model="form.<%==s item.Field%>">
            <Radio label="0">请自行编辑单选框</Radio>
            <Radio label="1">请自行编辑单选框</Radio>
          </RadioGroup>
          <%
          }
          %>
          <%
          if item.Type=="number" {
          %>
          <InputNumber v-model="form.<%==s item.Field%>" style="width:<%==s options.Width%>"></InputNumber>
          <%
          }
          %>
          <%
          if item.Type=="date"{
          %>
          <DatePicker type="date" v-model="form.<%==s item.Field%>" clearable style="width:<%==s options.Width%>"></DatePicker>
          <%
          }
          %>
          <%
            if item.Type=="daterange" {
          %>
          <DatePicker type="daterange" v-model="form.<%==s item.Field%>" clearable style="width:<%==s options.Width%>"></DatePicker>
          <%
          }
          %>
          <%
          if item.Type=="time" {
          %>
          <TimePicker type="time" v-model="form.<%==s item.Field%>" clearable style="width:<%==s options.Width%>"></TimePicker>
          <%
          }
          %>
          <%
          if item.Type=="area" {
          %>
          <al-cascader v-model="form.<%==s item.Field%>" data-type="code" level="<%==s item.Level%>" style="width:<%==s options.Width%>"/>
          <%
          }
          %>
          <%
          if item.Type=="slider" {
          %>
          <Slider v-model="form.<%==s item.Field%>" style="width:<%==s options.Width%>"></Slider>
          <%
          }
          %>
          <%
          if item.Type=="upload" {
          %>
          <upload-pic-input v-model="form.<%==s item.Field%>" style="width:<%==s options.Width%>"></upload-pic-input>
          <%
          }
          %>
          <%
          if item.Type=="uploadThumb" {
          %>
          <uploadThumb v-model="form.<%==s item.Field%>" multiple style="width:<%==s options.Width%>"></uploadThumb>
          <%
          }
          %>
          <%
          if item.Type=="wangEditor" && !options.ToQuill {
          %>
          <wangEditor v-model="form.<%==s item.Field%>" style="width:<%==s options.WangEditorWidth%>"></wangEditor>
          <%
          }
          %>
          <%
          if item.Type=="quill"||(item.Type=="wangEditor" && options.ToQuill){
          %>
          <quill v-model="form.<%==s item.Field%>" style="width:<%==s options.QuillWidth%>"></quill>
          <%
          }
          %>
          <%
          if item.Type=="password" {
          %>
          <password v-model="form.<%==s item.Field%>" style="width:<%==s options.Width%>"></password>
          <%
          }
          %>
        </FormItem>
        <%
          }
        }
        %>
      </Form>
      <div slot="footer">
        <Button type="text" @click="modalVisible=false">取消</Button>
        <Button type="primary" :loading="submitLoading" @click="handleSubmit">提交</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
<%
if options.Api {
%>
// 根据你的实际请求api.js位置路径修改
import { get<%==s options.ApiName%>List, add<%==s options.ApiName%>, edit<%==s options.ApiName%>, delete<%==s options.ApiName%> } from "@/api/index";
import moment from "moment";
<%
}
%>
<%
if options.Upload {
%>
import uploadPicInput from "@/views/my-components/xboot/upload-pic-input";
<%
}
%>
<%
if options.UploadThumb {
%>
import uploadThumb from "@/views/my-components/xboot/upload-pic-thumb";
<%
}
%>
<%
if options.WangEditor && !options.ToQuill {
%>
import wangEditor from "@/views/my-components/xboot/editor";
<%
}
%>
<%
if options.Quill||options.ToQuill {
%>
import quill from "@/views/my-components/xboot/quill";
<%
}
%>
<%
if options.Password {
%>
import password from "@/views/my-components/xboot/set-password";
<%
}
%>
export default {
  name: "<%==s options.VueName%>",
  components: {
    <%
    if options.Upload {
    %>
    uploadPicInput,
    <%
    }
    %>
    <%
    if options.UploadThumb {
    %>
    uploadThumb,
    <%
    }
    %>
    <%
    if options.WangEditor && !options.ToQuill {
    %>
    wangEditor,
    <%
    }
    %>
    <%
    if options.Quill||options.ToQuill {
    %>
    quill,
    <%
    }
    %>
    <%
    if options.Password {
    %>
    password,
    <%
    }
    %>
  },
  data() {
    return {
      <% if options.SearchSize>0 { %>
      openSearch: true, // 显示搜索
      <% }%>
      openTip: true, // 显示提示
      loading: true, // 表单加载状态
      modalType: 0, // 添加或编辑标识
      modalVisible: false, // 添加或编辑显示
      modalTitle: "", // 添加或编辑标题
      <% if options.HideSearch { %>
      drop: false,
      dropDownContent: "展开",
      dropDownIcon: "ios-arrow-down",
      <% } %>
      searchForm: { // 搜索框初始化对象
        pageNumber: 1, // 当前页数
        pageSize: 10, // 页面大小
        sort: "<%==s options.DefaultSort%>", // 默认排序字段
        order: "<%==s options.DefaultSortType%>", // 默认排序方式
        <% if options.DaterangeSearch { %>
        startDate: "", // 起始时间
        endDate: "" // 终止时间
        <% } %>
      },
      <% if options.DaterangeSearch { %>
      selectDate: null,
      <% } %>
      form: { // 添加或编辑表单对象初始化数据
        <%
        for _, item := range fields {
          if item.Editable {
        %>
        <% if item.Type=="switch" { %>
        <%==s item.Field%>: true,
        <% }else if item.Type=="number"|| item.Type=="slider" { %>
        <%==s item.Field%>: 0,
        <% }else if item.Type=="area"|| item.Type=="uploadThumb" { %>
        <%==s item.Field%>: [],
        <% }else{ %>
        <%==s item.Field%>: "",
        <% } %>
        <%
          }
        }
        %>
      },
      // 表单验证规则
      formValidate: {
        <%
        for _, item := range fields {
          if item.Editable&&item.Validate {
        %>
        <% if item.Type=="daterange"|| item.Type=="area"|| item.Type=="uploadThumb" { %>
        <%==s item.Field%>: [{ type: "array", required: true, message: "不能为空", trigger: "blur" }],
        <% }else if item.Type=="date" { %>
        <%==s item.Field%>: [{ required: true, message: "不能为空", trigger: "blur", pattern: /.+/ }],
        <% }else if item.Type=="number"|| item.Type=="slider" { %>
        <%==s item.Field%>: [{ type: "number", required: true, message: "不能为空", trigger: "blur" }],
        <% }else if item.Type=="switch" { %>
        <%==s item.Field%>: [{ type: "boolean", required: true, message: "不能为空", trigger: "blur" }],
        <% }else{ %>
        <%==s item.Field%>: [{ required: true, message: "不能为空", trigger: "blur" }],
        <% } %>
        <%
          }
        }
        %>
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
        <%
        for _, item := range fields {
          if item.TableShow {
        %>
        {
          title: "<%==s item.Name%>",
          key: "<%==s item.Field%>",
          minWidth: 120,
          <%
          if item.Sortable {
          %>
          sortable: true,
          <%
          }else{
          %>
          sortable: false,
          <%
          }
          %>
          <%
          if item.DefaultSort {
          %>
          sortType: "<%==s item.DefaultSortType%>"
          <%
          }
          %>
        },
        <%
          }
        }
        %>
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
    <% if options.SearchSize>0 { %>
    handleSearch() {
      this.searchForm.pageNumber = 1;
      this.searchForm.pageSize = 10;
      this.getDataList();
    },
    <% } %>
    handleReset() {
      this.$refs.searchForm.resetFields();
      this.searchForm.pageNumber = 1;
      this.searchForm.pageSize = 10;
      <% if options.DaterangeSearch { %>
      this.selectDate = null;
      this.searchForm.startDate = "";
      this.searchForm.endDate = "";
      <% } %>
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
    <% if options.DaterangeSearch { %>
    selectDateRange(v) {
      if (v) {
        this.searchForm.startDate = moment(v[0]).format("YYYY-MM-DD HH:mm:ss");
        this.searchForm.endDate = moment(v[1]).format("YYYY-MM-DD HH:mm:ss");
      }
    },
    <% } %>
    <% if options.HideSearch { %>
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
    <% } %>
    getDataList() {
      this.loading = true;
      <%
      if options.Api{
      %>
      get<%==s options.ApiName%>List(this.searchForm).then(res => {
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
      <%
      } else {
      %>
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
      <%
      }
      %>
    },
    handleSubmit() {
      this.$refs.form.validate(valid => {
        if (valid) {
        <%for _, item := range fields {
            if item.Type == "date" {%>
          this.form.<%==s item.Field%> = moment(this.form.<%==s item.Field%>).format("YYYY-MM-DD HH:mm:ss");
        <%}}%>
          this.submitLoading = true;
          if (this.modalType === 0) {
            // 添加 避免编辑后传入id等数据 记得删除
            delete this.form.id;
            <%
            if options.Api {%>
            add<%==s options.ApiName%>(this.form).then(res => {
              this.submitLoading = false;
              if (res.success) {
                this.$Message.success("操作成功");
                this.getDataList();
                this.modalVisible = false;
              }
            });
            <%
            } else {
            %>
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
            <%
            }
            %>
          } else {
            // 编辑
            <%
            if options.Api {
            %>
            edit<%==s options.ApiName%>(this.form).then(res => {
              this.submitLoading = false;
              if (res.success) {
                this.$Message.success("操作成功");
                this.getDataList();
                this.modalVisible = false;
              }
            });
            <%
            } else {
            %>
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
            <%
            }
            %>
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
          <%
          if options.Api {
          %>
          delete<%==s options.ApiName%>(v.id).then(res => {
            this.$Modal.remove();
            if (res.success) {
              this.$Message.success("操作成功");
              this.getDataList();
            }
          });
          <%
          } else {
          %>
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
          <%
          }
          %>
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
          <%
          if options.Api {
          %>
          delete<%==s options.ApiName%>(ids).then(res => {
            this.$Modal.remove();
            if (res.success) {
              this.$Message.success("操作成功");
              this.clearSelectAll();
              this.getDataList();
            }
          });
          <%
          } else {
          %>
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
          <%
          }
          %>
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
</style>