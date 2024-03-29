package v1

import (
	"gf-react-admin-server/app/model"
	"gf-react-admin-server/app/service"
	"gf-react-admin-server/library/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
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
		if _, ok := err.(*gvalid.Error); ok {
			// 参数校验失败
			response.FailCodeAndMsgExit(r, response.ParamValidErr, err.Error())
		}
		response.FailMsgExit(r, err.Error())
	}
	if err := service.Role.CreateRole(createRoleReq); err != nil {
		response.FailCodeAndMsgExit(r, gerror.Code(err), err.Error())
	}
	response.SuccessMsgExit(r, "创建角色成功")
}

// @summary 编辑角色接口
// @tags 角色服务
// @produce json
// @param entity body model.EditRoleReq true "编辑角色参数"
// @router /v1/role [PUT]
func (ra *roleApi) EditRole(r *ghttp.Request) {
	var editRoleReq = &model.EditRoleReq{}
	if err := r.ParseForm(editRoleReq); err != nil {
		if _, ok := err.(*gvalid.Error); ok {
			// 参数校验失败
			response.FailCodeAndMsgExit(r, response.ParamValidErr, err.Error())
		}
		response.FailMsgExit(r, err.Error())
	}
	if err := service.Role.EditRole(editRoleReq); err != nil {
		response.FailCodeAndMsgExit(r, gerror.Code(err), err.Error())
	}
	response.SuccessMsgExit(r, "编辑角色信息成功")
}

// @summary 角色分页列表
// @tags 角色服务
// @produce json
// @param entity body model.RoleQueryReq true "角色列表搜索参数"
// @router /v1/role [GET]
func (ra *roleApi) RolePageList(r *ghttp.Request) {
	var roleQueryReq = &model.RoleQueryReq{}
	if err := r.Parse(roleQueryReq); err != nil {
		if _, ok := err.(*gvalid.Error); ok {
			// 参数校验失败
			response.FailCodeAndMsgExit(r, response.ParamValidErr, err.Error())
		}
		response.FailMsgExit(r, err.Error())
	}

	roles, err := service.Role.RolePageList(roleQueryReq)
	if err != nil {
		response.FailCodeAndMsgExit(r, gerror.Code(err), err.Error())
	}
	response.SuccessDataExit(r, roles)
}

// @summary 角色详情
// @tags 角色服务
// @produce json
// @param id path int true "角色id"
// @router /v1/role/{id} [GET]
func (ra *roleApi) RoleById(r *ghttp.Request) {
	id := r.Get("id")
	if err := gvalid.Check(id, "required|integer|min:1", "角色id为空|参数类型非法|角色id长度非法"); err != nil {
		response.FailMsgExit(r, err.Error())
	}
	role, err := service.Role.RoleById(gconv.Int64(id))
	if err != nil {
		response.FailMsgExit(r, err.Error())
	}
	response.SuccessData(r, role)
}
