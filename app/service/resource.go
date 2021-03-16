package service

import (
	"gf-vue3-admin-server/app/dao"
	"gf-vue3-admin-server/app/model"
	"gf-vue3-admin-server/library/common"
	"gf-vue3-admin-server/library/response"
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
