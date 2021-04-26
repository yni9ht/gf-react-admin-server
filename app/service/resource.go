package service

import (
	"gf-react-admin-server/app/dao"
	"gf-react-admin-server/app/model"
	"gf-react-admin-server/library/common"
	"gf-react-admin-server/library/response"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
)

type resourceService struct{}

var (
	Resource = &resourceService{}
)

// CreateResource create resource service method
func (r *resourceService) CreateResource(req *model.CreateResourceReq) error {
	// 判断是否存在相同资源
	result, err := dao.Resource.FindByNameAndType(req.Name, req.Type)
	if err != nil {
		return err
	}
	if result != nil {
		return gerror.NewCode(response.DataExist, "该资源已存在，请勿重复创建。")
	}

	resource := &model.Resource{}
	err = gconv.Struct(req, resource)
	if err != nil {
		return err
	}

	// 生成id
	id, err := common.Id.GenerateUUID()
	if err != nil {
		return err
	}
	resource.Id = id

	// 创建资源
	_, err = dao.Resource.Insert(resource)
	if err != nil {
		return err
	}

	return nil
}

// EditResource 编辑资源信息
func (r *resourceService) EditResource(req *model.EditResourceReq) error {
	// 判断是否存在相同资源
	result, err := dao.Resource.FindByNameAndType(req.Name, req.Type)
	if err != nil {
		return err
	}
	if result != nil && result.Id != req.Id {
		return gerror.NewCode(response.DataExist, "资源名已存在，请修改资源名。")
	}

	// 更新
	_, err = dao.Resource.Update(req, dao.Resource.Columns.Id, req.Id)
	if err != nil {
		return err
	}
	return nil
}

// GetResourceTree 获取资源树形结构
func (r *resourceService) GetResourceTree() (trees []common.TreeNode, err error) {
	resources := make([]*model.Resource, 0)
	err = dao.Resource.Order("sn ASC").Structs(&resources)
	if err != nil {
		return nil, err
	}

	trees = common.GenerateTree(model.GetResourceSlice(resources))
	return
}

// GetMenuByUserId 获取用户具有的资源列表
func (r *resourceService) GetMenuByUserId(userId int) ([]*model.Resource, error) {
	if userId <= 0 {
		err := gerror.New("userId不能为空")
		return nil, err
	}

	userRelations := ([]model.Relation)(nil)
	err := dao.Relation.Where(dao.Relation.Columns.UserId, userId).Structs(&userRelations)
	if len(userRelations) <= 0 {
		return nil, err
	}

	roleIds := make([]int64, 0, len(userRelations))
	for _, v := range userRelations {
		roleIds = append(roleIds, v.RoleId)
	}

	roleResources := ([]model.RoleResource)(nil)
	err = dao.RoleResource.Where(dao.RoleResource.Columns.RoleId+" IN (?)", roleIds).Structs(&roleResources)
	if len(roleResources) <= 0 {
		return nil, err
	}

	resourceIds := make([]int64, 0, len(roleResources))
	for _, v := range roleResources {
		resourceIds = append(resourceIds, v.ResourceId)
	}

	resources := make([]*model.Resource, 0)
	err = dao.Resource.Where("type = ?", "menu").WherePri(resourceIds).Order("sn ASC").Structs(&resources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

// GetMenuTreeByUserId 根据用户id获取改用户拥有的菜单资源
func (r *resourceService) GetMenuTreeByUserId(userId int) (trees []common.TreeNode, err error) {
	resources, err := r.GetMenuByUserId(userId)
	if err != nil {
		return nil, err
	}

	trees = common.GenerateTree(model.GetResourceSlice(resources))
	return
}
