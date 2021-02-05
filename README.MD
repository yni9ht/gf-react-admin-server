# gf-vue3-admin-server

## 简介
该项目为`gf-vue3-admin`项目对应后端项目，采用`Golang` + [GoFrame](https://github.com/gogf/gf) 进行开发，支持JWT鉴权、casbin权限控制。

## 主要技术
- Golang
- [GoFrame](https://github.com/gogf/gf)
- JWT([gf-jwt](https://github.com/gogf/gf-jwt))
- [Casbin](https://github.com/casbin/casbin)
- MySQL
- Redis

## 项目细节
### 统一返回体
详情见项目`library`包
1. `resposne.go`文件，该文件中定义了统一的返回结构体如下：
```go
type ResultRes struct {
    Code    int         `json:"code"` // response code, default: success-200 error-500 响应码
    Message string      `json:"msg"`  // response message 响应信息，若返回的是错误码，则此处对应相关错误信息
    Data    interface{} `json:"data"` // response result data 返回数据
}
```
主要包含响应码、响应信息和相关数据，当响应码为200则表示服务正常，其余响应码对应相关错误，并会有对应的响应信息进行返回。
还定义了一些功能方法方便业务方进行调用返回，例如`Success(r *ghttp.Request)`、`SuccessExit(r *ghttp.Request)`，其中带有exit后缀的方法均表示会在返回response后退出当前HTTP执行函数

2. `page.go`定义了公共的分页返回参数结构体，并提供相关创建函数，结构体如下
```go
// 分页统一返回体
type PageResult struct {
    PageNum  int         `json:"pageNum"`  // 分页页数
    PageSize int         `json:"pageSize"` // 分页大小
    Total    int         `json:"total"`    // 总记录数
    Records  interface{} `json:"records"`  // 对应数据
}
```
