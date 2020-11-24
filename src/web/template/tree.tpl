<%!
import (
  "github.com/anden007/afocus-godf/src/web/view_model/generator"
)

type TreeOption struct{
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
  EditWidth string
  Span int
}
%>
<%: func Tree(fields []generator.Field, firstTwo []generator.Field, rest []generator.Field,options TreeOption, buffer *bytes.Buffer) %>
<template>
  <div class="search">
    <Card>
      <Row class="operation">
        <Button @click="add" type="primary" icon="md-add">添加子节点</Button>
        <Button @click="addRoot" icon="md-add">添加一级节点</Button>
        <Button @click="delAll" icon="md-trash">批量删除</Button>
        <Button @click="getParentList" icon="md-refresh">刷新</Button>
        <i-switch v-model="strict" size="large" style="margin-left:5px">
          <span slot="open">级联</span>
          <span slot="close">单选</span>
        </i-switch>
      </Row>
      <Row type="flex" justify="start">
        <Col span="6">
          <Alert show-icon>
            当前选择编辑：
            <span class="select-title">{{editTitle}}</span>
            <a class="select-clear" v-if="form.id" @click="cancelEdit">取消选择</a>
          </Alert>
          <Input
            v-model="searchKey"
            suffix="ios-search"
            @on-change="search"
            placeholder="输入节点名搜索"
            clearable
          />
          <div class="tree-bar" :style="{maxHeight: maxHeight}">
            <Tree
              ref="tree"
              :data="data"
              :load-data="loadData"
              show-checkbox
              @on-check-change="changeSelect"
              @on-select-change="selectTree"
              :check-strictly="!strict"
            ></Tree>
            <Spin size="large" fix v-if="loading"></Spin>
          </div>
        </Col>
        <Col span="<%==i  options.Span%>" style="margin-left:10px">
          <Form ref="form" :model="form" :label-width="100" :rules="formValidate" <% if options.RowNum>1{ %>inline<% } %>>
            <FormItem label="上级节点" prop="parentTitle">
              <div style="display:flex;width:<%==s  options.Width%>">
                <Input v-model="form.parentTitle" readonly style="margin-right:10px;"/>
                <Poptip transfer trigger="click" placement="right-start" title="选择上级部门" width="250">
                  <Button icon="md-list">选择部门</Button>
                  <div slot="content" style="position:relative;min-height:5vh">
                    <Tree :data="dataEdit" :load-data="loadData" @on-select-change="selectTreeEdit"></Tree>
                    <Spin size="large" fix v-if="loadingEdit"></Spin>
                  </div>
                </Poptip>
              </div>
            </FormItem>
            <FormItem label="名称" prop="title">
              <Input v-model="form.title" style="width:<%==s  options.Width%>"/>
            </FormItem>
            <%
            for _, item := range fields {
              if item.Editable {
            %>
            <FormItem label="<%==s  item.Name%>" prop="<%==s  item.Field%>" <% if options.RowNum>1&&(item.Type=="switch"||item.Type=="radio"){ %>style="width:<%==s  options.ItemWidth%>"<% } %> <% if item.Type=="wangEditor"||item.Type=="quill"{ %>class="form-<%==s  item.Type%>"<%}%>>
              <%
              if item.Type=="text" {
              %>
              <Input v-model="form.<%==s  item.Field%>" style="width:<%==s  options.EditWidth%>"/>
              <%
              }
              %>
              <%
              if item.Type=="select" {
              %>
              <Select v-model="form.<%==s  item.Field%>" style="width:<%==s  options.EditWidth%>">
                <Option value="0">请自行编辑下拉菜单</Option>
              </Select>
              <%
              }
              %>
              <%
              if item.Type=="number" {
              %>
              <InputNumber v-model="form.<%==s  item.Field%>" style="width:<%==s  options.EditWidth%>"></InputNumber>
              <%
              }
              %>
              <%
              if item.Type=="switch" {
              %>
              <i-switch v-model="form.<%==s  item.Field%>"></i-switch>
              <%
              }
              %>
              <%
              if item.Type=="radio" {
              %>
              <RadioGroup v-model="form.<%==s  item.Field%>">
                <Radio label="0">请自行编辑单选框</Radio>
                <Radio label="1">请自行编辑单选框</Radio>
              </RadioGroup>
              <%
              }
              %>
              <%
              if item.Type=="date" {
              %>
              <DatePicker type="date" v-model="form.<%==s  item.Field%>" style="width:<%==s  options.EditWidth%>"></DatePicker>
              <%
              }
              %>
              <%
                if item.Type=="daterange" {
              %>
              <DatePicker type="daterange" v-model="form.<%==s  item.Field%>" style="width:<%==s  options.EditWidth%>"></DatePicker>
              <%
              }
              %>
              <%
              if item.Type=="time" {
              %>
              <TimePicker type="time" v-model="form.<%==s  item.Field%>" style="width:<%==s  options.EditWidth%>"></TimePicker>
              <%
              }
              %>
              <%
              if item.Type=="area" {
              %>
              <al-cascader v-model="form.<%==s  item.Field%>" data-type="code" level="<%==s  item.Level%>" style="width:<%==s  options.EditWidth%>"/>
              <%
              }
              %>
              <%
              if item.Type=="slider" {
              %>
              <Slider v-model="form.<%==s  item.Field%>" style="width:<%==s  options.EditWidth%>"></Slider>
              <%
              }
              %>
              <%
              if item.Type=="upload" {
              %>
              <upload-pic-input v-model="form.<%==s  item.Field%>" style="width:<%==s  options.EditWidth%>"></upload-pic-input>
              <%
              }
              %>
              <%
              if item.Type=="uploadThumb" {
              %>
              <uploadThumb v-model="form.<%==s  item.Field%>" multiple style="width:<%==s  options.Width%>"></uploadThumb>
              <%
              }
              %>
              <%
              if item.Type=="wangEditor"&&!options.ToQuill {
              %>
              <wangEditor v-model="form.<%==s  item.Field%>" style="width:<%==s  options.WangEditorWidth%>"></wangEditor>
              <%
              }
              %>
              <%
              if item.Type=="quill"||(item.Type=="wangEditor"&&options.ToQuill){
              %>
              <quill v-model="form.<%==s  item.Field%>" style="width:<%==s  options.QuillWidth%>"></quill>
              <%
              }
              %>
              <%
              if item.Type=="password" {
              %>
              <password v-model="form.<%==s  item.Field%>" style="width:<%==s  options.Width%>"></password>
              <%
              }
              %>
            </FormItem>
            <%
              }
            }
            %>
            <FormItem label="排序值" prop="sortOrder" style="width:<%==s  options.ItemWidth%>">
              <Poptip trigger="hover" placement="right" content="值越小越靠前，支持小数">
                <InputNumber :max="1000" :min="0" v-model="form.sortOrder"></InputNumber>
              </Poptip>
            </FormItem>
            <br>
            <Form-item>
              <Button
                @click="submitEdit"
                :loading="submitLoading"
                type="primary"
                icon="ios-create-outline"
                style="margin-right:5px"
              >修改并保存</Button>
              <Button @click="handleReset">重置</Button>
            </Form-item>
          </Form>
        </Col>
      </Row>
    </Card>

    <Modal :title="modalTitle" v-model="modalVisible" :mask-closable="false" :width="<%==s  options.ModalWidth%>">
      <Form ref="formAdd" :model="formAdd" :label-width="100" :rules="formValidate" label-position="left" <% if(options.RowNum>1){ %>inline<% } %>>
        <div v-if="showParent">
          <FormItem label="上级节点：">{{form.title}}</FormItem>
        </div>
        <FormItem label="名称" prop="title">
          <Input v-model="formAdd.title" style="width:<%==s  options.Width%>"/>
        </FormItem>
        <%
        for _, item := range fields {
          if item.Editable {
        %>
        <FormItem label="<%==s  item.Name%>" prop="<%==s  item.Field%>" <% if options.RowNum>1&&(item.Type=="number"||item.Type=="switch"||item.Type=="radio"){ %>style="width:<%==s  options.ItemWidth%>"<% } %> <% if item.Type=="wangEditor"||item.Type=="quill" { %>class="form-<%==s  item.Type%>"<%}%>>
          <%
          if item.Type=="text" {
          %>
          <Input v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.Width%>"/>
          <%
          }
          %>
          <%
          if item.Type=="select" {
          %>
          <Select v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.Width%>">
            <Option value="0">请自行编辑下拉菜单</Option>
          </Select>
          <%
          }
          %>
          <%
          if item.Type=="number" {
          %>
          <InputNumber v-model="formAdd.<%==s  item.Field%>"></InputNumber>
          <%
          }
          %>
          <%
          if item.Type=="switch" {
          %>
          <i-switch v-model="formAdd.<%==s  item.Field%>"></i-switch>
          <%
          }
          %>
          <%
          if item.Type=="radio" {
          %>
          <RadioGroup v-model="formAdd.<%==s  item.Field%>">
            <Radio label="0">请自行编辑单选框</Radio>
            <Radio label="1">请自行编辑单选框</Radio>
          </RadioGroup>
          <%
          }
          %>
          <%
          if item.Type=="date" {
          %>
          <DatePicker type="date" v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.Width%>"></DatePicker>
          <%
          }
          %>
          <%
            if item.Type=="daterange" {
          %>
          <DatePicker type="daterange" v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.Width%>"></DatePicker>
          <%
          }
          %>
          <%
          if item.Type=="time" {
          %>
          <TimePicker type="time" v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.Width%>"></TimePicker>
          <%
          }
          %>
          <%
          if item.Type=="area" {
          %>
          <al-cascader v-model="formAdd.<%==s  item.Field%>" data-type="code" level="<%==s  item.Level%>" style="width:<%==s  options.Width%>"/>
          <%
          }
          %>
          <%
          if item.Type=="slider" {
          %>
          <Slider v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.Width%>"></Slider>
          <%
          }
          %>
          <%
          if item.Type=="upload" {
          %>
          <upload-pic-input v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.Width%>"></upload-pic-input>
          <%
          }
          %>
          <%
          if item.Type=="uploadThumb" {
          %>
          <uploadThumb v-model="formAdd.<%==s  item.Field%>" multiple style="width:<%==s  options.Width%>"></uploadThumb>
          <%
          }
          %>
          <%
          if item.Type=="wangEditor"&&!options.ToQuill {
          %>
          <wangEditor id="editor2" v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.WangEditorWidth%>"></wangEditor>
          <%
          }
          %>
          <%
          if item.Type=="quill"||(item.Type=="wangEditor"&& options.ToQuill ){
          %>
          <quill id="editor2" v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.QuillWidth%>"></quill>
          <%
          }
          %>
          <%
          if item.Type=="password" {
          %>
          <password v-model="formAdd.<%==s  item.Field%>" style="width:<%==s  options.Width%>"></password>
          <%
          }
          %>
        </FormItem>
        <%
          }
        }
        %>
        <FormItem label="排序值" prop="sortOrder" style="width:<%==s  options.ItemWidth%>">
          <Poptip trigger="hover" placement="right" content="值越小越靠前，支持小数">
            <InputNumber :max="1000" :min="0" v-model="formAdd.sortOrder"></InputNumber>
          </Poptip>
        </FormItem>
      </Form>
      <div slot="footer">
        <Button type="text" @click="modalVisible=false">取消</Button>
        <Button type="primary" :loading="submitLoading" @click="submitAdd">提交</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
<%
if options.Api {
%>
// 根据你的实际请求api.js位置路径修改
import { init<%==s  options.ApiName%>, load<%==s  options.ApiName%>, add<%==s  options.ApiName%>, edit<%==s  options.ApiName%>, delete<%==s  options.ApiName%>, search<%==s  options.ApiName%> } from "@/api/index";
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
if options.WangEditor&&!options.ToQuill {
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
  name: "<%==s  options.VueName%>",
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
    if options.WangEditor&&!options.ToQuill {
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
      maxHeight: "500px",
      strict: true,
      loading: false, // 树加载状态
      loadingEdit: false, // 编辑上级树加载状态
      modalVisible: false, // 添加显示
      selectList: [], // 多选数据
      selectCount: 0, // 多选计数
      showParent: false, // 显示上级标识
      modalTitle: "", // 添加标题
      editTitle: "", // 编辑节点名称
      searchKey: "", // 搜索树
      form: {
        // 编辑对象初始化数据
        id: "",
        title: "",
        parentId: "",
        parentTitle: "",
        sortOrder: 0,
        <%
        for _, item := range fields {
          if item.Editable {
        %>
        <% if item.Type=="switch" { %>
        <%==s  item.Field%>: true,
        <% }else if item.Type=="number"||item.Type=="slider" { %>
        <%==s  item.Field%>: 0,
        <% }else if item.Type=="area"||item.Type=="uploadThumb" { %>
        <%==s  item.Field%>: [],
        <% }else{ %>
        <%==s  item.Field%>: "",
        <% } %>
        <%
          }
        }
        %>
      },
      formAdd: {
        // 添加对象初始化数据
      },
      formValidate: {
        // 表单验证规则
        title: [{ required: true, message: "不能为空", trigger: "blur" }],
        sortOrder: [
          {
            required: true,
            type: "number",
            message: "排序值不能为空",
            trigger: "blur"
          }
        ],
        <%
        for _, item := range fields {
          if item.Editable&&item.Validate {
        %>
        <% if item.Type=="daterange"||item.Type=="area"||item.Type=="uploadThumb" { %>
        <%==s  item.Field%>: [{ type: "array", required: true, message: "不能为空", trigger: "blur" }],
        <% }else if item.Type=="date" { %>
        <%==s  item.Field%>: [{ required: true, message: "不能为空", trigger: "blur", pattern: /.+/ }],
        <% }else if item.Type=="number"||item.Type=="slider" { %>
        <%==s  item.Field%>: [{ type: "number", required: true, message: "不能为空", trigger: "blur" }],
        <% }else if item.Type=="switch" { %>
        <%==s  item.Field%>: [{ type: "boolean", required: true, message: "不能为空", trigger: "blur" }],
        <% }else{ %>
        <%==s  item.Field%>: [{ required: true, message: "不能为空", trigger: "blur" }],
        <% } %>
        <%
          }
        }
        %>
      },
      submitLoading: false,
      data: [],
      dataEdit: []
    };
  },
  methods: {
    init() {
      // 初始化一级节点
      this.getParentList();
      // 初始化一级节点为编辑上级节点使用
      this.getParentListEdit();
    },
    getParentList() {
      <%
      if options.Api {
      %>
      this.loading = true;
      init<%==s  options.ApiName%>().then(res => {
        this.loading = false;
        if (res.success) {
          res.result.forEach(function(e) {
            if (e.isParent) {
              e.loading = false;
              e.children = [];
            }
          });
          this.data = res.result;
        }
      });
      <%
      } else {
      %>
      // this.loading = true;
      // this.getRequest("一级数据请求路径，如/tree/getByParentId/0").then(res => {
      //   this.loading = false;
      //   if (res.success) {
      //     res.result.forEach(function(e) {
      //       if (e.isParent) {
      //         e.loading = false;
      //         e.children = [];
      //       }
      //     });
      //     this.data = res.result;
      //   }
      // });
      // 模拟请求成功
      this.data = [
      ];
      <%
      }
      %>
    },
    getParentListEdit() {
      <%
      if options.Api {
      %>
      this.loadingEdit = true;
      init<%==s  options.ApiName%>().then(res => {
        this.loadingEdit = false;
        if (res.success) {
          res.result.forEach(function(e) {
            if (e.isParent) {
              e.loading = false;
              e.children = [];
            }
          });
          // 头部加入一级
          let first = {
            id: "0",
            title: "一级节点"
          };
          res.result.unshift(first);
          this.dataEdit = res.result;
        }
      });
      <%
      } else {
      %>
      // this.loadingEdit = true;
      // this.getRequest("/tree/getByParentId/0").then(res => {
      //   this.loadingEdit = false;
      //   if (res.success) {
      //     res.result.forEach(function(e) {
      //       if (e.isParent) {
      //         e.loading = false;
      //         e.children = [];
      //       }
      //     });
      //     // 头部加入一级
      //     let first = {
      //       id: "0",
      //       title: "一级节点"
      //     };
      //     res.result.unshift(first);
      //     this.dataEdit = res.result;
      //   }
      // });
      // 模拟请求成功
      this.dataEdit = [
      ];
      <%
      }
      %>
    },
    loadData(item, callback) {
      <%
      if options.Api {
      %>
      load<%==s  options.ApiName%>(item.id).then(res => {
        if (res.success) {
          res.result.forEach(function(e) {
            if (e.isParent) {
              e.loading = false;
              e.children = [];
            }
          });
          callback(res.result);
        }
      });
      <%
      } else {
      %>
      // 异步加载树子节点数据
      // this.getRequest("请求路径，如/tree/getByParentId/" + item.id).then(res => {
      //   if (res.success) {
      //     res.result.forEach(function(e) {
      //       if (e.isParent) {
      //         e.loading = false;
      //         e.children = [];
      //       }
      //     });
      //     callback(res.result);
      //   }
      // });
      <%
      }
      %>
    },
    search() {
      // 搜索树
      if (this.searchKey) {
        <%
        if options.Api {
        %>
        this.loading = true;
        search<%==s  options.ApiName%>({ title: this.searchKey %>).then(res => {
          this.loading = false;
          if (res.success) {
            this.data = res.result;
          }
        });
        <%
        } else {
        %>
        // 模拟请求
        // this.loading = true;
        // this.getRequest("搜索请求路径", { title: this.searchKey }).then(res => {
        //   this.loading = false;
        //   if (res.success) {
        //     this.data = res.result;
        //   }
        // });
        // 模拟请求成功
        this.data = [
        ];
        <%
        }
        %>
      } else {
        // 为空重新加载
        this.getParentList();
      }
    },
    selectTree(v) {
      if (v.length > 0) {
        // 转换null为""
        for (let attr in v[0]) {
          if (v[0][attr] === null) {
            v[0][attr] = "";
          }
        }
        let str = JSON.stringify(v[0]);
        let data = JSON.parse(str);
        this.form = data;
        this.editTitle = data.title;
      } else {
        this.cancelEdit();
      }
    },
    cancelEdit() {
      let data = this.$refs.tree.getSelectedNodes()[0];
      if (data) {
        data.selected = false;
      }
      this.$refs.form.resetFields();
      this.form.id = "";
      this.editTitle = "";
    },
    selectTreeEdit(v) {
      if (v.length > 0) {
        // 转换null为""
        for (let attr in v[0]) {
          if (v[0][attr] === null) {
            v[0][attr] = "";
          }
        }
        let str = JSON.stringify(v[0]);
        let data = JSON.parse(str);
        this.form.parentId = data.id;
        this.form.parentTitle = data.title;
      }
    },
    handleReset() {
      this.$refs.form.resetFields();
      this.form.status = 0;
    },
    submitEdit() {
      this.$refs.form.validate(valid => {
        if (valid) {
        <%for _, item := range fields {
            if item.Type == "date" {%>
          this.form.<%==s item.Field%> = moment(this.form.<%==s item.Field%>).format("YYYY-MM-DD HH:mm:ss");
        <%}}%>
          if (!this.form.id) {
            this.$Message.warning("请先点击选择要修改的节点");
            return;
          }
          this.submitLoading = true;
          <%
          if options.Api {
          %>
          edit<%==s  options.ApiName%>(this.form).then(res => {
          this.submitLoading = false;
          if (res.success) {
            this.$Message.success("编辑成功");
              this.init();
              this.modalVisible = false;
            }
          });
          <%
          } else {
          %>
          // this.postRequest("请求路径，如/tree/edit", this.form).then(res => {
          //   this.submitLoading = false;
          //   if (res.success) {
          //     this.$Message.success("编辑成功");
          //     this.init();
          //     this.modalVisible = false;
          //   }
          // });
          // 模拟成功
          this.submitLoading = false;
          this.$Message.success("编辑成功");
          this.modalVisible = false;
          <%
          }
          %>
        }
      });
    },
    submitAdd() {
      this.$refs.formAdd.validate(valid => {
        if (valid) {
          this.submitLoading = true;
          <%
          if options.Api {
          %>
          add<%==s  options.ApiName%>(this.formAdd).then(res => {
            this.submitLoading = false;
            if (res.success) {
              this.$Message.success("添加成功");
              this.init();
              this.modalVisible = false;
            }
          });
          <%
          } else {
          %>
          // this.postRequest("请求路径，如/tree/add", this.formAdd).then(res => {
          //   this.submitLoading = false;
          //   if (res.success) {
          //     this.$Message.success("添加成功");
          //     this.init();
          //     this.modalVisible = false;
          //   }
          // });
          // 模拟成功
          this.submitLoading = false;
          this.$Message.success("添加成功");
          this.modalVisible = false;
          <%
          }
          %>
        }
      });
    },
    add() {
      if (this.form.id == "" || this.form.id == null) {
        this.$Message.warning("请先点击选择一个节点");
        return;
      }
      this.modalTitle = "添加子节点";
      this.showParent = true;
      this.formAdd = {
        parentId: this.form.id,
        sortOrder: 0,
        <%
        for _, item := range fields {
          if item.Editable {
        %>
        <% if item.Type=="switch" { %>
        <%==s  item.Field%>: true,
        <% }else if item.Type=="number"||item.Type=="slider" { %>
        <%==s  item.Field%>: 0,
        <% }else if item.Type=="area"||item.Type=="uploadThumb" { %>
        <%==s  item.Field%>: [],
        <% }else{ %>
        <%==s  item.Field%>: "",
        <% } %>
        <%
          }
        }
        %>
      };
      this.modalVisible = true;
    },
    addRoot() {
      this.modalTitle = "添加一级节点";
      this.showParent = false;
      this.formAdd = {
        parentId: 0,
        sortOrder: 0,
        title: "",
        <%
        for _, item := range fields {
          if item.Editable {
        %>
        <% if item.Type=="switch" { %>
        <%==s  item.Field%>: true,
        <% }else if item.Type=="number"||item.Type=="slider" { %>
        <%==s  item.Field%>: 0,
        <% }else if item.Type=="area" { %>
        <%==s  item.Field%>: [],
        <% }else{ %>
        <%==s  item.Field%>: "",
        <% } %>
        <%
          }
        }
        %>
      };
      this.modalVisible = true;
    },
    changeSelect(v) {
      this.selectCount = v.length;
      this.selectList = v;
    },
    delAll() {
      if (this.selectCount <= 0) {
        this.$Message.warning("您还未勾选要删除的数据");
        return;
      }
      this.$Modal.confirm({
        title: "确认删除",
        content: "您确认要删除所选的 " + this.selectCount + " 条数据及其下级所有数据?",
        loading: true,
        onOk: () => {
          let ids = "";
          this.selectList.forEach(function(e) {
            ids += e.id + ",";
          });
          ids = ids.substring(0, ids.length - 1);
           <%
          if options.Api {
          %>
          delete<%==s  options.ApiName%>(ids).then(res => {
            this.$Modal.remove();
            if (res.success) {
              this.$Message.success("删除成功");
              this.selectList = [];
              this.selectCount = 0;
              this.cancelEdit();
              this.init();
            }
          });
          <%
          } else {
          %>
          // this.deleteRequest("请求路径，如/tree/delByIds/" + ids).then(res => {
          //   this.$Modal.remove();
          //   if (res.success) {
          //     this.$Message.success("删除成功");
          //     this.selectList = [];
          //     this.selectCount = 0;
          //     this.cancelEdit();
          //     this.init();
          //   }
          // });
          // 模拟成功
          this.$Modal.remove();
          this.$Message.success("删除成功");
          this.selectList = [];
          this.selectCount = 0;
          this.cancelEdit();
          <%
          }
          %>
        }
      });
    }
  },
  mounted() {
    // 计算高度
    let height = document.documentElement.clientHeight;
    this.maxHeight = Number(height-287) + "px";
    this.init();
  }
};
</script>
<style lang="less">
// 建议引入通用样式 具体路径自行修改 可删除下面样式代码
// @import "../../../styles/tree-common.less";
.search {
    .operation {
        margin-bottom: 2vh;
    }
    .select-title {
        font-weight: 600;
        color: #40a9ff;
    }
    .select-clear {
        margin-left: 10px;
    }
}

.tree-bar {
    overflow: auto;
    margin-top: 5px;
    position: relative;
    min-height: 80px;
}

.tree-bar::-webkit-scrollbar {
    width: 6px;
    height: 6px;
}

.tree-bar::-webkit-scrollbar-thumb {
    border-radius: 4px;
    -webkit-box-shadow: inset 0 0 2px #d1d1d1;
    background: #e4e4e4;
}
</style>