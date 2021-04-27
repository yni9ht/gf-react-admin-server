package middleware

import (
	"gf-react-admin-server/app/model"
	"gf-react-admin-server/app/service"
	"gf-react-admin-server/library/response"
	"github.com/gogf/gcache-adapter/adapter"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
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
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
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
		PayloadFunc:   PayloadFunc,
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
// @summary 登录
// @tags 用户服务
// @produce json
// @param entity body model.UserLoginReq true "登录参数"
// @router /v1/user/login [POST]
func Authenticator(r *ghttp.Request) (interface{}, error) {
	var (
		loginReg *model.UserLoginReq
	)
	if err := r.Parse(&loginReg); err != nil {
		g.Log().Errorf("登录失败，err:[%v]", err)
		response.FailMsgExit(r, err.Error())
	}
	loginRes, err := service.User.Login(loginReg)
	return gconv.Map(&loginRes), err
}

// PayloadFunc customer info to jwt. 放入自定义信息到jwt的claims中
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

//TODO Authorization check user auth role by IdentityHandler func. 根据IdentityHandler方法返回的信息检查用户权限，返回false则检查失败
func Authorization(data interface{}, r *ghttp.Request) bool {
	return true
}

// Unauthorized custom unauthorized result. 自定义无权限返回方法
func Unauthorized(r *ghttp.Request, code int, message string) {
	g.Log().Warningf("访问失败，错误码：[%v], 错误信息:[%v], 请求路径：[%v],请求方法：[%v], router-method:[%v], router-uri:[%v], router-regNames [%v], router-regRule:[%v]",
		code, message,
		r.RequestURI, r.Method, r.Router.Method, r.Router.Uri, r.Router.RegNames, r.Router.RegRule)
	if message == "" {
		message = errorMsg
	}
	response.FailCodeAndMsgExit(r, code, message)
}

// LoginResponse custom callback login response. 自定义登录返回方法
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.RestResultExit(r, code, "", &loginRes{
		Token:  token,
		Expire: expire,
	})
}

// JwtAuth request head token valid. 检查请求中的token合法性
func JwtAuth(r *ghttp.Request) {
	JwtMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}
