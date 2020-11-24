<%!
import (
    "github.com/anden007/afocus-godf/src/web/view_model/generator"
)

type TableIndexOption struct{
    SearchSize int
    HideSearch bool
    ApiName string
    DaterangeSearch bool
    Api bool
    DefaultSort string
    DefaultSortType string
}
%>
<%: func TableIndex(fields []generator.Field, firstTwo []generator.Field, rest []generator.Field, options TableIndexOption, buffer *bytes.Buffer) %>
<template>
    <div class="search">
        <add v-if="currView=='add'" @close="currView='index'" @submited="submited" />
        <edit v-if="currView=='edit'" @close="currView='index'" @submited="submited" :data="formData" />
        <Card v-show="currView=='index'">
            <%
            if options.SearchSize>0&&!options.HideSearch {
            %>
            <Row <% if options.SearchSize>0{ %>v-show="openSearch"<% } %> @keydown.enter.native="handleSearch">
            <Form ref="searchForm" :model="searchForm" inline :label-width="70">
                <%
                for _,item := range fields{
                if item.Searchable {
                %>
                <%
                if(item.SearchType=="text"){
                %>
                <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                    <Input type="text" v-model="searchForm.<%==s  item.Field%>" placeholder="请输入<%==s  item.Name%>" clearable style="width: 200px"/>
                </Form-item>
                <%
                }
                %>
                <%
                if item.SearchType=="select" {
                %>
                <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                    <Select v-model="searchForm.<%==s  item.Field%>" placeholder="请选择" clearable style="width: 200px">
                        <Option value="0">请自行编辑下拉菜单</Option>
                    </Select>
                </Form-item>
                <%
                }
                %>
                <%
                if item.SearchType=="date" {
                %>
                <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                    <DatePicker type="date" v-model="searchForm.<%==s  item.Field%>" placeholder="请选择" clearable style="width: 200px"></DatePicker>
                </Form-item>
                <%
                }
                %>
                <%
                if item.SearchType=="daterange" {
                %>
                <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                    <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
                </Form-item>
                <%
                }
                %>
                <%
                if item.SearchType=="area" {
                %>
                <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                    <al-cascader v-model="searchForm.<%==s  item.Field%>" data-type="code" level="<%==s  item.SearchLevel%>" style="width:200px"/>
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
            if options.SearchSize > 0 && options.HideSearch {
            %>
            <Row @keydown.enter.native="handleSearch">
                <Form ref="searchForm" :model="searchForm" inline :label-width="70" class="search-form">
                    <%
                    for _,item := range firstTwo {
                    %>
                    <%
                    if item.SearchType=="text" {
                    %>
                    <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                        <Input type="text" v-model="searchForm.<%==s  item.Field%>" placeholder="请输入<%==s  item.Name%>" clearable style="width: 200px"/>
                    </Form-item>
                    <%
                    }
                    %>
                    <%
                    if item.SearchType=="select" {
                    %>
                    <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                        <Select v-model="searchForm.<%==s  item.Field%>" placeholder="请选择" clearable style="width: 200px">
                            <Option value="0">请自行编辑下拉菜单</Option>
                        </Select>
                    </Form-item>
                    <%
                    }
                    %>
                    <%
                    if item.SearchType=="date" {
                    %>
                    <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                        <DatePicker type="date" v-model="searchForm.<%==s  item.Field%>" placeholder="请选择" clearable style="width: 200px"></DatePicker>
                    </Form-item>
                    <%
                    }
                    %>
                    <%
                    if item.SearchType=="daterange" {
                    %>
                    <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                        <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
                    </Form-item>
                    <%
                    }
                    %>
                    <%
                    if item.SearchType=="area" {
                    %>
                    <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                        <al-cascader v-model="searchForm.<%==s  item.Field%>" data-type="code" level="<%==s  item.SearchLevel%>" style="width:200px"/>
                    </Form-item>
                    <%
                    }
                    %>
                    <%
                    }
                    %>
                    <span v-if="drop">
            <%
            for _, item := range rest{
            %>
              <%
              if item.SearchType=="text" {
              %>
              <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                <Input type="text" v-model="searchForm.<%==s  item.Field%>" placeholder="请输入<%==s  item.Name%>" clearable style="width: 200px"/>
              </Form-item>
              <%
              }
              %>
              <%
              if item.SearchType=="select" {
              %>
              <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                <Select v-model="searchForm.<%==s  item.Field%>" placeholder="请选择" clearable style="width: 200px">
                  <Option value="0">请自行编辑下拉菜单</Option>
                </Select>
              </Form-item>
              <%
              }
              %>
              <%
              if item.SearchType=="date" {
              %>
              <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                <DatePicker type="date" v-model="searchForm.<%==s  item.Field%>" placeholder="请选择" clearable style="width: 200px"></DatePicker>
              </Form-item>
              <%
              }
              %>
              <%
              if item.SearchType=="daterange" {
              %>
              <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                <DatePicker v-model="selectDate" type="daterange" format="yyyy-MM-dd" clearable @on-change="selectDateRange" placeholder="选择起始时间" style="width: 200px"></DatePicker>
              </Form-item>
              <%
              }
              %>
              <%
              if item.SearchType=="area" {
              %>
              <Form-item label="<%==s  item.Name%>" prop="<%==s  item.Field%>">
                <al-cascader v-model="searchForm.<%==s  item.Field%>" data-type="code" level="<%==s  item.SearchLevel%>" style="width:200px"/>
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
                <Button type="dashed" @click="openTip=!openTip">{{openTip ? "关闭提示" : "开启提示"}}</Button>
            </Row>
            <Row v-show="openTip">
                <Alert show-icon>
                    已选择
                    <span class="select-count">{{selectCount}}</span> 项
                    <a class="select-clear" @click="clearSelectAll">清空</a>
                </Alert>
            </Row>
            <Row>
                <Table
                        :loading="loading"
                        border
                        :columns="columns"
                        :data="data"
                        ref="table"
                        sortable="custom"
                        @on-sort-change="changeSort"
                        @on-selection-change="changeSelect"
                ></Table>
            </Row>
            <Row type="flex" justify="end" class="page">
                <Page
                        :current="searchForm.pageNumber"
                        :total="total"
                        :page-size="searchForm.pageSize"
                        @on-change="changePage"
                        @on-page-size-change="changePageSize"
                        :page-size-opts="[10,20,50]"
                        size="small"
                        show-total
                        show-elevator
                        show-sizer
                ></Page>
            </Row>
        </Card>
    </div>
</template>

<script>
    <%
    if options.Api {
    %>
// 根据你的实际请求api.js位置路径修改
    import { get<%==s options.ApiName%>List, delete<%==s options.ApiName%> } from "@/api/index";
    import moment from "moment";
    <%
    }
    %>
    // 根据你的实际添加编辑组件位置路径修改
    import add from "./add.vue";
    import edit from "./edit.vue";
    export default {
        name: "single-window",
        components: {
            add,
            edit
        },
        data() {
            return {
                <% if options.SearchSize>0 { %>
                openSearch: true, // 显示搜索
            <% }%>
            openTip: true, // 显示提示
                formData: {},
            currView: "index",
                loading: true, // 表单加载状态
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
            for _, item := range fields{
                if item.TableShow {
                %>
                    {
                        title: "<%==s  item.Name%>",
                            key: "<%==s  item.Field%>",
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
                            sortType: "<%==s  item.DefaultSortType%>"
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
                pageNumber: 1, // 当前页数
                pageSize: 10, // 页面大小
                total: 0 // 表单数据总数
        };
        },
        methods: {
                init() {
                    this.getDataList();
                },
                submited() {
                    this.currView = "index";
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
    clearSelectAll() {
        this.$refs.table.selectAll(false);
    },
    changeSelect(e) {
        this.selectList = e;
        this.selectCount = e.length;
    },
    getDataList() {
        this.loading = true;
    <%
        if options.Api {
        %>
            get <%==s options.ApiName%>List(this.searchForm).then(res => {
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
    add() {
        this.currView = "add";
    },
    edit(v) {
        // 转换null为""
        for (let attr in v) {
            if (v[attr] == null) {
                v[attr] = "";
            }
        }
        let str = JSON.stringify(v);
        let data = JSON.parse(str);
        this.formData = data;
        this.currView = "edit";
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