package service

import (
	"gf-react-admin-server/app/model"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestCreateRelation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		req := &model.SaveRelationReq{
			UserId:  1,
			RoleIds: []int{1, 2, 4},
		}
		t.Assert(Relation.SaveRelation(req), nil)
	})
}

func TestCreateRelationErr1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		req := &model.SaveRelationReq{
			UserId: 1,
		}
		if err := Relation.SaveRelation(req); err != nil {
			t.Log(err)
		} else {
			t.Error("测试失败")
		}
	})
}

func TestCreateRelationErr2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		req := &model.SaveRelationReq{
			RoleIds: []int{1},
		}
		if err := Relation.SaveRelation(req); err != nil {
			t.Log(err)
		} else {
			t.Error("测试失败")
		}
	})
}
