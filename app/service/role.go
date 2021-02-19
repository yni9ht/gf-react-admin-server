package service

import (
	"gf-vue3-admin-server/app/dao"
	"gf-vue3-admin-server/app/model"
	"gf-vue3-admin-server/library/common"
	"gf-vue3-admin-server/library/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
)

type roleService struct{}

var (
	// Role 用于外部调用
	Role = &roleService{}
)

// CreateRole 创建角色
func (r *roleService) CreateRole(req *model.CreateRoleReq) error {
	// 判断是否已经存在该角色
	role, err := dao.Role.FindByNameOrAlias(req.RoleName, req.Alias)
	if err != nil {
		return err
	}
	if role != nil {
		return gerror.NewCode(response.DataExist, "该角色名或别名已存在，请勿重复创建。")
	}

	role = &model.Role{}
	err = gconv.Struct(req, role)
	if err != nil {
		return err
	}

	// 生成id
	id, err := common.GenerateUUID()
	if err != nil {
		return err
	}
	role.Id = int64(id)
	// 创建角色
	_, err = dao.Role.Insert(role)
	if err != nil {
		return err
	}

	return nil
}
