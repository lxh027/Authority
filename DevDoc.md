# 权限管理系统开发文档

## 后端

### 综述

#### 技术栈

- go1.15.2 

- gin

- gorm

#### 环境安装

##### 安装go环境

##### 开启go mod

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

##### 在GoLand中开启go mod

##### 在命令行中运行main.go

```shell
go run main.go
```

#### 项目配置

##### 项目目录

    ├─app
    │  ├─api
    │  │  ├─controller
    │  │  │  └─IndexController.go
    │  │  ├─model
    │  │  │  └─IndexModel.go
    │  │  └─validate
    │  │     └─IndexValidate.go
    │  └─common
    ├─config
    ├─db_server
    ├─log
    ├─routes
    ├─server
    ├─go.mod
    ├─main.go
##### 数据库配置

在`config/database.go`中

```go
func GetDbConfig() map[string]interface{} {

	// init db config
	dbConfig := make(map[string]interface{})

	dbConfig["hostname"] 	= "localhost"
	dbConfig["port"] 		= "3306"
	dbConfig["database"] 	= ""
	dbConfig["username"] 	= "root"
	dbConfig["password"] 	= ""
	dbConfig["charset"]		= "utf8"
	dbConfig["parseTime"]	= "True"

	return dbConfig
}
```

##### 端口配置

在`config/server`中，改变端口号

```go
serverConfig["port"] 	= "5000"
```

则运行后访问端口`localhost:5000`即能访问

#### 开发

`app`中可分模块进行开发，每个模块内分别有三部分组成：`controller`、`model`、`validate`

##### Model

`Model`为连接数据库的操作，使用`gorm`库构建，文档地址：[gorm](http://gorm.book.jasperxu.com)

##### Validate

`Validate`模块为对数据进行验证，比如对权限模块进行验证，只需要在`validate`文件夹中创建`AuthorityValidate.go`，并编写如下内容：

```go
package validate

import "OnlineJudge/app/common"

var AuthorityValidate common.Validator

func init()  {
	rules := map[string]string{
		"id": "required",
	}

	scenes := map[string][]string {
		"find": {"id"},
	}
	AuthorityValidate.Rules = rules
	AuthorityValidate.Scenes = scenes
}
```

其中，`rules`map中填写需要验证的字段，`scenes`中填写对于一个特定的验证场景需要验证哪些字段。

则在使用验证时，只需要在对应`controller`中使用：

```go
func Test(c *gin.Context)  {
	var authorityValidate = validate.AuthorityValidate

	if res, err := userValidate.Validate(c, "find"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}
}
```

具体`validate`文档地址：[go validate](https://gitee.com/inhere/validate)

##### Controller

`Controller`中为定义的接口，可以接收网络请求，一个基本的`Controller`请求如下：

```go
func Register(c *gin.Context)  {
	var userModel = model.User{}
	var userValidate = validate.UserValidate

    // 数据验证
	if res, err := userValidate.Validate(c, "register"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	password, passwordCheck := c.PostForm("password"), c.PostForm("password_check")

	if password != passwordCheck {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "两次密码输入不一致", ""))
	}

    // 数据绑定
	var userJson model.User
	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	userJson.Password = common.GetMd5(userJson.Password)

	res := userModel.AddUser(userJson)

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
```

主要分为处理请求、验证数据、获取数据、访问`Model`获取数据、返回接口几个部分。[gin](https://www.kancloud.cn/shuangdeyu/gin_book/949411)

##### Session

需要使用Session时，首先`import "github.com/gin-contrib/sessions"`

在`controller`中，使用首先初始化一个`session`对象

```go
session := sessions.Default(c)
```

存储`session`

```go
session.Set("data", data)
session.Save()
```

获取`session`

```go
res := session.Get("user_id")
```

若`session`不存在，则`res = nil`

##### 路由

在`routes/router.go`中，使用分组方式划分路由，如想设置一个`api/user/getAllUser`的接口

```go
func Routes(router *gin.Engine)  {

	// api
	api := router.Group("/api")
	{
        user := api.Group("/user")
        {
            user.POST("/getAllUser", controller.GetAllUser)
        }
	}

}
```

其中`controller.GetAllUser`为`controller/UserController`下的一个方法

#### 接口需求

|       接口名称        |                功能 难度                 |
| :-------------------: | :--------------------------------------: |
|    `user/register`    |                用户注册 0                |
|     `user/login`      |                用户登录 1                |
|     `user/logout`     |                用户登出 1                |
|   `user/checkLogin`   | 检查登录状态，若为管理员，返回权限列表 2 |
|    `user/addRole`     |              添加用户角色 0              |
|   `user/deleteRole`   |              删除用户角色 0              |
|   `user/getAllAuth`   |            获取用户权限列表 2            |
|    `user/getMenu`     |            获取用户后台列表 2            |
|    `role/addRole`     |                添加角色 0                |
|   `role/deleteRole`   |                删除角色 0                |
|    `role/editRole`    |                编辑角色 0                |
|  `role/addRoleAuth`   |              添加角色权限 0              |
| `role/deleteRoleAuth` |              删除角色权限 0              |
|    `auth/addAuth`     |                添加权限 0                |
|   `auth/deleteAuth`   |                删除权限 0                |
|    `auth/editAuth`    |                编辑权限 0                |

## 前端

使用`layuimini`开发使用

文档地址：

- `layui`:[layui](https://www.layui.com/doc/)
- `layuimini`:[layuimini](http://layuimini.99php.cn/docs/index.html)



|       页面        |       功能       |
| :---------------: | :--------------: |
| `auth/index.html` |     权限列表     |
|  `auth/add.html`  |     添加权限     |
| `auth/edit.html`  |     编辑权限     |
| `role/index.html` |     角色列表     |
|  `role/add.html`  |     添加角色     |
| `role/edit.html`  |     编辑角色     |
| `role/auth.html`  | 角色所有权限列表 |
| `user/index.html` |     用户列表     |
| `user/role.html`  | 用户所有角色列表 |

