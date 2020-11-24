<%!
import (
  "github.com/anden007/afocus-godf/src/web/view_model/generator"
)

type AddOption struct{
    RowNum int
    ItemWidth string
    Width string
    WangEditorWidth string
    QuillWidth string
    ApiName string
    Api bool
    Upload bool
    UploadThumb bool
    WangEditor bool
    Quill bool
    Password bool
}
%>
<%: func Add(fields []generator.Field,options AddOption, buffer *bytes.Buffer) %>
<template>
  <div>
    <Card>
      <div slot="title">
        <div class="edit-head">
          <a @click="close" class="back-title">
            <Icon type="ios-arrow-back" />返回
          </a>
          <div class="head-name">添加</div>
          <span></span>
          <a @click="close" class="window-close">
            <Icon type="ios-close" size="31" class="ivu-icon-ios-close" />
          </a>
        </div>
      </div>
      <Form ref="form" :model="form" :label-width="100" :rules="formValidate" label-position="left" <% if options.RowNum>1 { %>inline<% } %>>
          <%
          for _, item := range fields{
            if item.Editable {
          %>
          <FormItem label="<%==s item.Name%>" prop="<%==s item.Field%>" <% if options.RowNum>1&&(item.Type=="switch"||item.Type=="radio"){ %>style="width:<%==s options.ItemWidth%>"<% } %> <% if item.Type=="wangEditor"||item.Type=="quill"{ %>class="form-<%==s item.Type%>"<% } %>>
            <%
            if item.Type=="text" {
            %>
            <Input v-model="form.<%==s item.Field%>" clearable style="width:<%==s options.Width%>"/>
            <%
            }
            %>
            <%
            if item.Type=="select"{
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
            if item.Type=="radio"{
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
            if item.Type=="date" {
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
            if item.Type=="wangEditor" {
            %>
            <wangEditor v-model="form.<%==s item.Field%>" style="width:<%==s options.WangEditorWidth%>"></wangEditor>
            <%
            }
            %>
            <%
            if item.Type=="quill" {
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
          <Form-item class="br">
            <Button
              @click="handleSubmit"
              :loading="submitLoading"
              type="primary"
            >提交并保存</Button>
            <Button @click="handleReset">重置</Button>
            <Button type="dashed" @click="close">关闭</Button>
          </Form-item>
        </Form>
    </Card>
  </div>
</template>

<script>
<%
if options.Api {
%>
// 根据你的实际请求api.js位置路径修改
import { add<%==s options.ApiName%> } from "@/api/index";
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
if options.WangEditor {
%>
import wangEditor from "@/views/my-components/xboot/editor";
<%
}
%>
<%
if options.Quill {
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
  name: "add",
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
      if options.WangEditor {
      %>
      wangEditor,
      <%
      }
      %>
      <%
      if options.Quill {
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
      submitLoading: false, // 表单提交状态
      form: { // 添加或编辑表单对象初始化数据
        <%
        for _,item := range fields {
          if item.Editable {
        %>
        <% if item.Type=="switch" { %>
        <%==s item.Field%>: true,
        <% }else if item.Type=="number"||item.Type=="slider" { %>
        <%==s item.Field%>: 0,
        <% }else if item.Type=="area"||item.Type=="uploadThumb" { %>
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
        for _,item := range fields {
          if item.Editable&&item.Validate {
        %>
        <% if item.Type=="daterange"||item.Type=="area"||item.Type=="uploadThumb" { %>
        <%==s item.Field%>: [{ type: "array", required: true, message: "不能为空", trigger: "blur" }],
        <% }else if item.Type=="date" { %>
        <%==s item.Field%>: [{ required: true, message: "不能为空", trigger: "blur", pattern: /.+/ }],
        <% }else if item.Type=="number"||item.Type=="slider" { %>
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
      }
    };
  },
  methods: {
    init() {},
    handleReset() {
      this.$refs.form.resetFields();
    },
    handleSubmit() {
      this.$refs.form.validate(valid => {
        if (valid) {
        <%for _, item := range fields {
            if item.Type == "date" {%>
          this.form.<%==s item.Field%> = moment(this.form.<%==s item.Field%>).format("YYYY-MM-DD HH:mm:ss");
        <%}}%>
          <%
          if options.Api {
          %>
            add<%==s options.ApiName%>(this.form).then(res => {
            this.submitLoading = false;
            if (res.success) {
              this.$Message.success("操作成功");
              this.submited();
            }
          });
          <%
          } else {
          %>
          // this.postRequest("请求路径", this.form).then(res => {
          //   this.submitLoading = false;
          //   if (res.success) {
          //     this.$Message.success("添加成功");
          //     this.submited();
          //   }
          // });
          // 模拟成功
          this.submitLoading = false;
          this.$Message.success("添加成功");
          this.submited();
          <%
          }
          %>
        }
      });
    },
    close() {
      this.$emit("close", true);
    },
    submited() {
      this.$emit("submited", true);
    }
  },
  mounted() {
    this.init();
  }
};
</script>
<style lang="less">
// 建议引入通用样式 具体路径自行修改 可删除下面样式代码
// @import "../../../styles/single-common.less";
.edit-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    position: relative;

    .back-title {
        color: #515a6e;
        display: flex;
        align-items: center;
    }

    .head-name {
        display: inline-block;
        height: 20px;
        line-height: 20px;
        font-size: 16px;
        color: #17233d;
        font-weight: 500;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }

    .window-close {
        z-index: 1;
        font-size: 12px;
        position: absolute;
        right: 0px;
        top: -5px;
        overflow: hidden;
        cursor: pointer;

        .ivu-icon-ios-close {
            color: #999;
            transition: color .2s ease;
        }
    }

    .window-close .ivu-icon-ios-close:hover {
        color: #444;
    }
}
</style>