package common

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

type Tee struct {
	Id       int
	Name     string
	ParentId int
}

func (r Tee) GetPrimKey() int {
	return r.Id
}

func (r Tee) GetParentPrimKey() int {
	return r.ParentId
}

func (r Tee) GetName() string {
	return r.Name
}

func (r Tee) GetData() interface{} {
	return r
}

func (r Tee) Root() bool {
	return r.ParentId == 0
}

func getData(datas []Tee) (iTrees []ITreeNode) {
	for _, v := range datas {
		iTrees = append(iTrees, v)
	}
	return
}

func TestCreateRelation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		datas := []Tee{
			{Id: 1, Name: "根节点1", ParentId: 0},
			{Id: 2, Name: "根节点2", ParentId: 0},
			{Id: 3, Name: "根节点3", ParentId: 0},
			{Id: 21, Name: "子节点1", ParentId: 2},
			{Id: 11, Name: "子节点2", ParentId: 1},
			{Id: 31, Name: "子节点3", ParentId: 3},
			{Id: 211, Name: "子节点4", ParentId: 21},
			{Id: 111, Name: "子节点5", ParentId: 11},
			{Id: 311, Name: "子节点6", ParentId: 31},
			{Id: 1111, Name: "子节点7", ParentId: 111},
		}
		tree := GenerateTree(getData(datas))
		by, _ := json.Marshal(tree)
		fmt.Print(string(by))
		t.AssertNE(tree, nil)
	})
}
