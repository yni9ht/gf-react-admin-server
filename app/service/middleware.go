package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Middleware = new(middlewareService)

type middlewareService struct{}

// 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *middlewareService) GetPath(r *ghttp.Request) {
	// 获取请求路径
	g.Log().Infof("请求路径：[%v],请求方法：[%v], router-method:[%v], router-uri:[%v], router-regNames [%v], router-regRule:[%v]",
		r.RequestURI, r.Method, r.Router.Method, r.Router.Uri, r.Router.RegNames, r.Router.RegRule)
	r.Middleware.Next()
}
