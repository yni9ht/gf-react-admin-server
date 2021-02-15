package router

import (
	"gf-vue3-admin-server/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(
			middleware.CORS,
			middleware.GetPath)
		group.POST("/user/login", middleware.JwtMiddleware.LoginHandler)
		group.GET("/user/logout", middleware.JwtMiddleware.LogoutHandler)
	})
}
