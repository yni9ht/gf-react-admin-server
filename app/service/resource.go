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
