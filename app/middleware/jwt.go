package middleware

import (
	"gf-vue3-admin-server/library/response"
	"github.com/gogf/gcache-adapter/adapter"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

type JwtConfig struct {
	ExpireTime    int    // 过期时间
	RefreshTime   int    // 刷新时间
	SignKey       string // 签名
	Realm         string // 域
	IdentityKey   string // 鉴权中用户唯一标识
	TokenLookup   string // token在请求中位置
	TokenHeadName string // token前缀
}

var (
	// gf-jwt middleware. gf-jwt自带的中间件
	JwtMiddleware *jwt.GfJWTMiddleware
	// jwt config. jwt配置
	JwtCfg   *JwtConfig
	errorMsg = "无访问权限"
)

type loginRes struct {
	token  string
	expire time.Time
}

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

//TODO Authenticator user login valid, and return userInfo. 进行登录校验，返回用户信息、错误
func Authenticator(r *ghttp.Request) (interface{}, error) {
	return nil, nil
}

//TODO Authorization check user auth role. 检查用户权限，返回false则检查失败。
func Authorization(data interface{}, r *ghttp.Request) bool {
	return true
}

// Unauthorized custom unauthorized result. 自定义无权限返回方法
func Unauthorized(r *ghttp.Request, code int, message string) {
	g.Log().Warningf("访问失败，错误码：[%v], 错误信息:[%v], 请求路径：[%v],请求方法：[%v], router-method:[%v], router-uri:[%v], router-regNames [%v], router-regRule:[%v]",
		code, message,
		r.RequestURI, r.Method, r.Router.Method, r.Router.Uri, r.Router.RegNames, r.Router.RegRule)
	response.FailCodeAndMsgExit(r, code, errorMsg)
}

// LoginResponse custom callback login response. 自定义登录返回方法
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.RestResultExit(r, code, "", &loginRes{
		token:  token,
		expire: expire,
	})
}

// JwtAuth request head token valid. 检查请求中的token合法性
func JwtAuth(r *ghttp.Request) {
	JwtMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}
