# gf-react-admin-server

## 简介
该项目为`gf-react-admin`项目对应后端项目，前端项目见[gf-react-admin](https://github.com/Riunshow/gf-react-admin)，采用`Golang` + [GoFrame](https://github.com/gogf/gf) 进行开发，支持JWT鉴权、casbin权限控制。

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
### JWT
本项目采用 [gf-jwt](https://github.com/gogf/gf-jwt) ，jwt的一些基础配置例如过期时间、签名key等都放在`config/config.toml`文件的中。jwt相关业务方发详情见`app/middleware/jwt.go`文件，该文件中的`init()`方法中初始化了
gf-jwt的相关配置，通过自定义的结构体对外进行调用。`Authenticator`、`Authorization`、
`Unauthorized`、`LoginResponse`分别对应了登录方法、鉴权方法、token校验失败方法、登录返回
方法，用来替换gf-jwt的对应默认方法。其他配置详情见gf-jwt的相关文档和案例。
```go
// init to gf-jwt config. 进行jwt初始化
func init() {
    // 初始化配置
    JwtCfg = new(JwtConfig)
    if err := g.Cfg().GetStruct("jwt", JwtCfg); err != nil {
        g.Log().Errorf("load jwt config fail, err: %v", err)
    }
    jwtMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
        Realm:         JwtCfg.Realm,
        Key:           []byte(JwtCfg.SignKey),
        Timeout:       time.Duration(JwtCfg.ExpireTime) * time.Minute,  // 过期时间
        MaxRefresh:    time.Duration(JwtCfg.RefreshTime) * time.Minute, // 刷新时间
        IdentityKey:   JwtCfg.IdentityKey,
        TokenLookup:   JwtCfg.TokenLookup,
        TokenHeadName: JwtCfg.TokenHeadName,
        TimeFunc:      time.Now,
        CacheAdapter:  adapter.NewRedis(g.Redis("cache")), // 采用redis替换默认的内存缓存
        Authenticator: Authenticator,
        Authorizator:  Authorization,
        Unauthorized:  Unauthorized,
        LoginResponse: LoginResponse,
    })
    if err != nil {
        g.Log().Errorf("init jwt middleware fail, error: %v", err)
    }
    JwtMiddleware = jwtMiddleware
}
```
