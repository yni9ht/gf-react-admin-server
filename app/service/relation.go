package service

import (
	"gf-vue3-admin-server/app/dao"
	"gf-vue3-admin-server/app/model"
	"gf-vue3-admin-server/library/common"
	"github.com/gogf/gf/util/gvalid"
)

type relationService struct{}

var (
	// Relation 供外部调用
	Relation = &relationService{}
)

// SaveRelation 保存用户关系
func (r *relationService) SaveRelation(req *model.SaveRelationReq) error {
	if err := gvalid.CheckStruct(&req, nil); err != nil {
		return err
	}

	// 删除旧关系
	deleteParam := map[string]interface{}{
		dao.Relation.Columns.UserId:                 req.UserId,
		dao.Relation.Columns.RoleId + " not in (?)": req.RoleIds,
	}
	_, _ = dao.Relation.Delete(deleteParam)

	for _, v := range req.RoleIds {
		// 根据用户id和需要新增的角色id查询是否已存在该关联关系
		queryParam := map[string]interface{}{
			dao.Relation.Columns.UserId: req.UserId,
			dao.Relation.Columns.RoleId: v,
		}
		count, err := dao.Relation.Where(queryParam).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		var relation = &model.Relation{
			UserId: int64(req.UserId),
			RoleId: int64(v),
		}
		relation.Id, err = common.Id.GenerateUUID()
		if err != nil {
			return err
		}
		// 创建
		_, err = dao.Relation.Insert(relation)
		if err != nil {
			return err
		}
	}

	return nil
}
