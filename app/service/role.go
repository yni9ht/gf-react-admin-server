package service

import (
	"gf-react-admin-server/app/dao"
	"gf-react-admin-server/app/model"
	"gf-react-admin-server/library/common"
	"gf-react-admin-server/library/response"
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
	id, err := common.Id.GenerateUUID()
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

// EditRole 编辑角色信息
func (r *roleService) EditRole(req *model.EditRoleReq) error {
	// 判断是否已经存在该角色
	role, err := dao.Role.FindByNameOrAlias(req.RoleName, req.Alias)
	if err != nil {
		return err
	}
	if role != nil && role.Id != req.Id {
		return gerror.NewCode(response.DataExist, "该角色名或别名已存在，请修改角色名或别名。")
	}

	err = gconv.Struct(req, role)
	if err != nil {
		return err
	}

	// 更新角色
	_, err = dao.Role.Update(role, dao.Role.Columns.Id, role.Id)
	return err
}

// RolePageList 获取角色分页列表
func (r *roleService) RolePageList(req *model.RoleQueryReq) (interface{}, error) {
	pageNo, pageSize := req.PageNo, req.PageSize
	total, err := dao.Role.Count()
	if err != nil {
		return nil, err
	}
	roles := make([]model.RolePageListRes, 0)
	err = dao.Role.Page(pageNo, pageSize).Structs(&roles)
	if err != nil {
		return nil, err
	}

	return &common.PageResult{
		PageNo:   pageNo,
		PageSize: pageSize,
		Total:    total,
		Records:  roles,
	}, nil
}

// RoleById 根据id获取角色信息
func (r *roleService) RoleById(id int64) (role *model.RoleInfoReq, err error) {
	err = dao.Role.WherePri(id).Struct(&role)
	return
}
