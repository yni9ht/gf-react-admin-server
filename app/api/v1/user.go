package v1

import (
	"gf-react-admin-server/app/service"
	"gf-react-admin-server/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

// apiUser is user api and func receiver
type userApi struct{}

var (
	// User to global call
	User = &userApi{}
)

// GetUserInfos
// @summary 用户详情
// @tags 用户服务
// @produce json
// @param id path int true "用户id"
// @router /v1/user/{id} [GET]
func (u *userApi) GetUserInfos(r *ghttp.Request) {
	id := r.Get("id")
	if err := gvalid.Check(id, "required|integer|min:1", "用户id为空|参数类型非法|用户id长度非法"); err != nil {
		response.FailMsgExit(r, err.Error())
	}
	userInfos, err := service.User.GetUserInfos(gconv.Int(id))
	if err != nil {
		response.FailMsgExit(r, err.Error())
	}
	response.SuccessData(r, userInfos)
}
