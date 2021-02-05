package router

import (
	"gf-vue3-admin-server/app/api/hello"
	"gf-vue3-admin-server/app/service"
	"gf-vue3-admin-server/library/common"
	"gf-vue3-admin-server/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.GetPath)
		group.ALL("/", hello.Hello)
		group.ALL("/result", func(r *ghttp.Request) {
			response.RestResult(r, response.OK, response.OkMessage, nil)
			r.Exit()
		})
		group.ALL("/resultExit", func(r *ghttp.Request) {
			response.RestResultExit(r, response.OK, response.OkMessage, nil)
		})
		group.ALL("/success", func(r *ghttp.Request) {
			response.Success(r)
			r.Exit()
		})
		group.ALL("/successExit", func(r *ghttp.Request) {
			response.SuccessExit(r)
		})
		group.ALL("/successMsg", func(r *ghttp.Request) {
			response.SuccessMsg(r, "这是一个自定义消息，手动退出。")
			r.Exit()
		})
		group.ALL("/successMsgExit", func(r *ghttp.Request) {
			response.SuccessMsgExit(r, "这是一个自定义消息，自动退出。")
		})
		group.ALL("/successData", func(r *ghttp.Request) {
			response.SuccessData(r, map[string]interface{}{
				"pageNum":  1,
				"pageSize": 10,
			})
			r.Exit()
		})
		group.ALL("/successDataExit", func(r *ghttp.Request) {
			response.SuccessDataExit(r, common.NewResult(3, 30, 100, []int{1, 23}))
		})
	})
}
