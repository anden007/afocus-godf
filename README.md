### 一、准备工作
由于系统内有部分库在Gitee(码云)内，要正常使用，请在系统环境变量中添加

`变量：GOPRIVATE 值：gitee.com,*.gitee.com`

### 安装开发期命令行工具

【必须安装】

`go get -u github.com/google/wire/cmd/wire` 

`go get -u github.com/ungerik/pkgreflect`

`go get -u github.com/go-bindata/go-bindata/...`

`go get -u github.com/swaggo/swag/cmd/swag`

---

### 二、gen脚本用法
#### 自动更新Go和Vue模板文件(不是必须，仅在模板发生变化后才需要运行)
`gen tpl`
#### 自动生成Go的CRUD文件
`gen model [model_name]`

model_name：模块名称,注意大小写，建议大驼峰方式命名
#### 按提示在main.go中进行模块注册和路由绑定即可！
#### 应用发布
由于框架采用bindata库将静态资源打包到了应用中，请在发布前一定要执行一次以下命令，否则会在生产环境运行出错！
 
`gen asset`

#### 使用swagger

按swagger规范要求，在controllers中做好相关方法的注释，然后使用以下命令进行初始化

`gen swag`

---
### 三、备忘录
#### 1、json序列化时，忽略空值字段的方法:
在字段tag中用omitempty关键字进行标记，即可当此字段为“0值”时不序列化，例如：

`gorm:"foreignkey:ParentId;association_foreignkey:Id" json:"children,omitempty"`
####json序列化和反序列化时忽略某个字段
在字段tag中用json:"-"即可，例如：

`gorm:"foreignkey:ParentId;association_foreignkey:Id" json:"-"`

#### 2、gorm中Raw和Scan搭配使用时，一定要注意字段名（或别名）一定要满足gorm的命名规则，否则无法正确获取到数据

#### 3、数据模型使用说明
##### tag标签使用方法
`json:"fieldName"` 表明该字段在被序列化成json格式时的key名称，fieldName需要自行替换，通常跟字段名一致（小驼峰命名法，例如：userName）

`gorm:"..."` gorm所用标签，请查阅官网教程 https://gorm.io/docs/models.html

##### 小技巧：
如果某个字段是gorm的关联对象（例如一对一、一对多、多对多），那么不建议将此标签设置成跟字段名一样，因为这样会造成页面添加和修改时由于无法正确解析关联字段对象而报错，建议在字段名前面加"_"，例如”_userName",这样就可以避免出错啦。例子如下：

```
type StoreItem struct {
 	Id          string    `gorm:"size:36;primary_key" json:"id"`
 	StoreId     string    `gorm:"size:36" json:"storeId"`
 	Store       *Store    `json:"store"`
 	BoxTypeId   string    `gorm:"size:36" json:"boxTypeId"`
 	BoxType     *BoxType  `json:"boxType"`
 	RemainCount int       `gorm:"default:0" json:"remainCount" caption:"库存数量"`
 	CreateTime  time.Time `json:"createTime"`
 }
```
#### 4、关于时区的陷阱
在go中如果直接使用`time.Parse`进行时间解析，解析出来的将会是UTC制式，而通常我们用来获取当前时间的`time.now()`返回的时间是当前系统设置的时区时间，如果直接用这两个时间做对比，将会有8小时误差！！！

解决方案：

请使用`time.ParseInLocation("2006-01-02 15:04:05", timeString, time.Local)`进行时间解析