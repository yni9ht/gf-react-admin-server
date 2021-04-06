package router

import (
	v1 "gf-react-admin-server/app/api/v1"
	"gf-react-admin-server/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	// RouteGroupPrefix 路由统一前缀
	RouteGroupPrefix = "/v1"
)

func init() {
	s := g.Server()
	// 跨域中间件
	s.Use(middleware.CORS)
	// 异常记录中间件
	s.Use(middleware.ErrorLog)
	s.Group(RouteGroupPrefix, func(group *ghttp.RouterGroup) {
		group.Middleware(
			middleware.GetPath)
		group.POST("/user/login", middleware.JwtMiddleware.LoginHandler)
		group.GET("/user/logout", middleware.JwtMiddleware.LogoutHandler)
	})

	// 角色服务
	s.Group(RouteGroupPrefix, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.JwtAuth)
		group.POST("/role", v1.Role.CreateRole)
		group.PUT("/role", v1.Role.EditRole)
		group.GET("/role", v1.Role.RolePageList)
		group.GET("/role/{id}", v1.Role.RoleById)
	})

	// 用户关系服务
	s.Group(RouteGroupPrefix, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.JwtAuth)
		group.PUT("/relation", v1.Relation.SaveRelation)
	})

	// 资源服务
	s.Group(RouteGroupPrefix, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.JwtAuth)
		group.POST("/resource", v1.Resource.CreateResource)
		group.PUT("/resource", v1.Resource.EditResource)
		group.GET("/resource/tree", v1.Resource.GetResourceTree)
	})
}
