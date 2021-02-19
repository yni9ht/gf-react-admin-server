package v1

import (
	"gf-vue3-admin-server/app/model"
	"gf-vue3-admin-server/app/service"
	"gf-vue3-admin-server/library/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
)

type roleApi struct{}

var (
	// Role 外部调用
	Role = &roleApi{}
)

// @summary 创建角色接口
// @tags 角色服务
// @produce json
// @param entity body model.CreateRoleReq true "创建角色参数"
// @router /v1/role [POST]
func (ra *roleApi) CreateRole(r *ghttp.Request) {
	var createRoleReq = &model.CreateRoleReq{}
	if err := r.Parse(createRoleReq); err != nil {
		response.FailMsgExit(r, err.Error())
	}
	if err := service.Role.CreateRole(createRoleReq); err != nil {
		response.FailCodeAndMsgExit(r, gerror.Code(err), err.Error())
	}
	response.SuccessMsgExit(r, "创建角色成功")
}