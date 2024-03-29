// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Role is the golang structure for table admin_role.
type Role struct {
	Id        int64       `orm:"id,primary" json:"id"`        //
	ParentId  string      `orm:"parent_id"  json:"parentId"`  // 父角色id
	RoleName  string      `orm:"role_name"  json:"roleName"`  // 角色名称
	Alias     string      `orm:"alias"      json:"alias"`     // 角色别名
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` // 更新时间
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 删除时间时间
}
