// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"gf-react-admin-server/app/model/internal"
	"gf-react-admin-server/library/common"
)

// Resource is the golang structure for table admin_resource.
type Resource internal.Resource

// Fill with you ideas below.
func (r *Resource) GetPrimKey() int {
	return int(r.Id)
}

func (r *Resource) GetParentPrimKey() int {
	return int(r.ParentId)
}

func (r *Resource) GetName() string {
	return r.Name
}

func (r *Resource) GetData() interface{} {
	return r
}

func (r *Resource) Root() bool {
	return r.ParentId == 0
}

// GetResourceSlice 将菜单资源slice转换为树形结构通用slice
func GetResourceSlice(data []*Resource) (iTrees []common.ITreeNode) {
	for _, v := range data {
		iTrees = append(iTrees, v)
	}
	return
}

// CreateResourceReq 新增资源参数
type CreateResourceReq struct {
	ParentId int64  `json:"parentId"`                                                       // 父资源id
	Name     string `json:"name" v:"name@required#资源名称不能为空"`                                // 资源名称
	Alias    string `json:"alias"`                                                          // 资源别称
	Url      string `json:"url"`                                                            // 资源路径
	Enable   bool   `json:"enable"`                                                         // 0:不显示，1显示
	Icon     string `json:"icon"`                                                           // 资源图标
	Type     string `json:"type" v:"type@required|in:menu,button,link#资源类型不能为空|请选择正确的资源类型"` // 资源类型：menu-菜单；button-按钮；link-链接
	Sn       int    `json:"sn"`                                                             // 排序
}

// EditResourceReq 编辑资源参数
type EditResourceReq struct {
	Id       int64  `json:"id" v:"id@required|min:1#待编辑资源id不能为空|待编辑资源id不能为空"`               // 主键id
	ParentId int64  `json:"parentId"`                                                       // 父资源id
	Name     string `json:"name" v:"name@required#资源名称不能为空"`                                // 资源名称
	Alias    string `json:"alias"`                                                          // 资源别称
	Url      string `json:"url"`                                                            // 资源路径
	Enable   bool   `json:"enable"`                                                         // 0:不显示，1显示
	Icon     string `json:"icon"`                                                           // 资源图标
	Type     string `json:"type" v:"name@required|in:menu,button,link#资源类型不能为空|请选择正确的资源类型"` // 资源类型：menu-菜单；button-按钮；link-链接
	Sn       int    `json:"sn"`                                                             // 排序
}

// ResourceTree 资源树形结构
type ResourceTree struct {
	Resource
	Child []ResourceTree `json:"child"` // 子资源
}

// ResourceRoleTree 角色对应的资源树形结构
type ResourceRoleTree struct {
	Resource
	Checked bool               `json:"checked"` // 是否选中
	Child   []ResourceRoleTree `json:"child"`   // 子资源
}

// ResourceReq 资源信息
type ResourceInfo struct {
	Id       int64  `json:"id"`       // 主键id
	ParentId int64  `json:"parentId"` // 父资源id
	Name     string `json:"name"`     // 资源名称
	Alias    string `json:"alias"`    // 资源别称
	Url      string `json:"url"`      // 资源路径
	Enable   bool   `json:"enable"`   // 0:不显示，1显示
	Icon     string `json:"icon"`     // 资源图标
	Type     string `json:"type"`     // 资源类型：menu-菜单；button-按钮；link-链接
	Sn       int    `json:"sn"`       // 排序
}
