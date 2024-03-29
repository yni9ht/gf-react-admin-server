// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// User is the golang structure for table admin_user.
type User struct {
	Id        int64       `orm:"id,primary" json:"id"`        //
	FullName  string      `orm:"full_name"  json:"fullName"`  // 姓名
	Account   string      `orm:"account"    json:"account"`   // 账号
	Password  string      `orm:"password"   json:"password"`  // 密码
	Email     string      `orm:"email"      json:"email"`     // 邮箱
	Mobile    string      `orm:"mobile"     json:"mobile"`    // 手机号码
	Wechat    string      `orm:"wechat"     json:"wechat"`    // 微信号
	Avatar    string      `orm:"avatar"     json:"avatar"`    // 头像
	Sex       int         `orm:"sex"        json:"sex"`       // 性别：0:未知，1:男，2:女
	Status    int         `orm:"status"     json:"status"`    // 0:禁用，1正常
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` // 更新时间
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 删除时间时间
}
