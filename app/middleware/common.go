package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 允许接口跨域请求中间件
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 异常记录中间件
func ErrorLog(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		g.Log().Error(err)
	}
}

func GetPath(r *ghttp.Request) {
	// 获取请求路径
	g.Log().Infof("请求路径：[%v],请求方法：[%v], router-method:[%v], router-uri:[%v], router-regNames [%v], router-regRule:[%v]",
		r.RequestURI, r.Method, r.Router.Method, r.Router.Uri, r.Router.RegNames, r.Router.RegRule)
	r.Middleware.Next()
}
