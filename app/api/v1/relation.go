package v1

import (
	"gf-react-admin-server/app/model"
	"gf-react-admin-server/app/service"
	"gf-react-admin-server/library/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type relationApi struct{}

var (
	Relation = &relationApi{}
)

// @summary 保存用户关系
// @tags 用户关系服务
// @produce json
// @param entity body model.SaveRelationReq true "保存用户关系参数"
// @router /v1/relation [PUT]
func (r *relationApi) SaveRelation(req *ghttp.Request) {
	saveRelationReq := &model.SaveRelationReq{}
	if err := req.Parse(saveRelationReq); err != nil {
		if _, ok := err.(*gvalid.Error); ok {
			// 参数校验失败
			response.FailCodeAndMsgExit(req, response.ParamValidErr, err.Error())
		}
		response.FailMsgExit(req, err.Error())
	}
	if err := service.Relation.SaveRelation(saveRelationReq); err != nil {
		response.FailCodeAndMsgExit(req, gerror.Code(err), err.Error())
	}
	response.SuccessMsgExit(req, "保存成功")
}
