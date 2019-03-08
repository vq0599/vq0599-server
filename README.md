# 基于Golang Gin的RESTful API实践
## 前言
职业生涯早期作为一名“打杂工程师”时接触过一些Go语言，后来就专注于前端了~
两年多了，试着拾起来~虽然没有改变外部环境的能力，但是希望自己有能力做得更好。这个时代，“伪全栈”已经成为每一位合格前端的标配。

## 项目结构
```bash
├── conf 
├── common 
├── middleware
├── util
├── models
├── controller
├── routers
├── main.go
└── README.md
```
- `conf`：配置文件。
- `common`： 和业务相关的一些公用方法和变量。（例：统一错误信息、统一返回格式、统一请求参数验证）
- `controller`：处理业务逻辑和数据。
- `modals`：和数据库交互。
- `util`：和业务无关的公用方法。（例：工具方法、library二次封装）
- `middleware`：API中间件。（如：权限验证）
- `routers`：路由设计。
- `main.go`：项目入口


## 项目构建思路
### API流程
```
request -> 路由 -> 鉴权 -> request参数处理（Controller） -> 数据库获取数据（Modal） -> response数据处理（Controller） -> response。
```
### 参数规则
遵循以下三点。如更新一篇文章的标题，需要文章id和标题title，则id通过URL传递，即`/api/articles/${id}`，title通过request payload传递。
- 资源唯一标识符ID，使用URL。
- 非唯一标识符参数，增加/修改数据，使用request payload。
- 限制条件，使用QueryString。
### 返回信息
- 统一的返回结构
  ```js
  {
    code: 'SUCCESS',
    data: null,
    msg: '成功'
  }
  ```
- 关于使用HTTP status code还是数据层面的`code`，一直是非常有争议的话题。我倒发表不了什么高见，只是对于写给自己用的API，我的原则是**系统层面的信息使用HTTP status code，业务层面的信息则使用具体数据code**。表示业务信息的`code`和`msg`定义成常量，`code`需要与HTTP status code有明显的不同或使用语义化更好的字符串，以保持直观的辨识度。例如：发布一篇文章：
  - 未登录：则使用HTTP status code（401）
  - 套餐不足，请购买套餐：则使用数据层面的业务code
- 数据结构不变性。例如获取文件的标签为空时，应该返回为空数组，而不是null。


## 路由设计
### 路由格式
API路由格式统一为：域名 + `/api`前缀 + ['模块名称'] + '资源名称' + '资源ID'。
```js
// 后台管理系统：v1版本 - admin模块 - id为10 - 文章资源
`${DOMAIN}/api/v1/admin/articles/10`
```
为了避免路由设计过长，有时会把“模块名称”省略。这个项目中，我统一把无权限验证的前台接口省略了模块名称。
```js
// 面向用户的前台：v1版本 - id为10 - 文章资源
`${DOMAIN}/api/v1/articles/10`
```
### 版本号
关于版本号，最早了到解这一概念是在原生客户端的场景。由于其无法热更新的特性，服务端在迭代功能时必须保证老版本客户端使用的接口不能受到影响。向下兼容本身就是一件比较痛苦的事情，版本号意义就在于解放了新旧代码的耦合。你不需要一遍一遍的打补丁，而是把旧代码扔到安静的角落，然后开始新的表演。

### HTTP方法
因为一些历史原因和一些“恶略的环境”，许多人习惯于只有`Get`和`Post`方法的HTTP请求。实际上HTTP协议一共有8种请求方法，分别都有不同的语义化场景[MDN戳我](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Methods)。我坚定的支持使用更加符合语义化规范的HTTP方法，就像我知道用勺子吃面也行，但是我选择了筷子。

## 鉴权体系
项目的鉴权系统是基于[JWT(JSON Web Tokens)](https://jwt.io/)构建的。JWT的原理比较简单，在用户登录成功时，会生成TOKEN返回给客户端。客户端发起API请求时携带TOKE即可。我们在middleware目录中写好了权限验证的中间件的`jwt.go`, 给需要验证的路由加上即可。
```go
apiAdmin.Use(middleware.Jwt())
{
  apiAdmin.POST("/articles", controller.AddArticle)
  // ...others
}
```
需要提及的是，单纯的JWT鉴权存在一个不可弥补的缺陷**无法服务端退出**，本质上就一种**以牺牲安全性为代价换取便捷性**的方案。所以你可能会看见许多基于JWT魔改的数据库通信方案（开始我也搞了，后来删除，实际意义不大~）。对于个人项目场景下，JWT显然是足够用的。

## 静态资源
一般在生产项目中，为了减小服务器的压力都会把静态资源放置在单独的服务器上。因此我倒腾出了之前购买腾讯云学生优惠套餐赠送的50G的COS存储空间，基于腾讯云提供的SDK腾讯云提供了COS存储的SDK[tencentcloud-sdk-go](https://github.com/tencentcloud/tencentcloud-sdk-go)实现了一个上传图片的API。文件的命名上使用了[shortid](https://github.com/teris-io/shortid)，保证每次上传的图片的名称都不重复。（上传同路径同名的文件时，腾讯云默认会把旧的文件替换了）

## 结语
实际上，我至今对GO语言的了解程度都十分粗浅。完成这个项目，也仅仅是过了一遍文档和各种“集百家之长”。期待未来能实践于生产环境中，期待更好~