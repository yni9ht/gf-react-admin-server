package v1

import (
	"gf-vue3-admin-server/app/model"
	"gf-vue3-admin-server/app/service"
	"gf-vue3-admin-server/library/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type resourceApi struct{}

var (
	Resource = &resourceApi{}
)

// @summary 新建资源信息
// @tags 资源管理
// @produce json
// @param entity body model.CreateResourceReq true "新建资源参数"
// @router /v1/resource [POST]
func (r *resourceApi) CreateResource(req *ghttp.Request) {
	createResource := &model.CreateResourceReq{}

	if err := req.Parse(createResource); err != nil {
		if _, ok := err.(*gvalid.Error); ok {
			// 参数校验失败
			response.FailCodeAndMsgExit(req, response.ParamValidErr, err.Error())
		}
		response.FailMsgExit(req, err.Error())
	}
	if err := service.Resource.CreateResource(createResource); err != nil {
		response.FailCodeAndMsgExit(req, gerror.Code(err), err.Error())
	}
	response.SuccessMsgExit(req, "新建成功")
}
