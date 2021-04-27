package service

import (
	"gf-react-admin-server/app/dao"
	"gf-react-admin-server/app/model"
	"gf-react-admin-server/library/response"
	"github.com/gogf/gf/errors/gerror"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
}

var (
	User = &userService{}
)

func (u *userService) Login(loginReq *model.UserLoginReq) (user *model.UserLoginRes, err error) {
	var queryParam = map[string]string{
		"account": loginReq.Account,
	}
	err = dao.User.Where(queryParam).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.NewCode(response.AccountExist, "账户不存在")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, gerror.NewCode(response.AccountValidErr, "账户或密码错误")
	}
	return user, nil
}

// GetUserInfos 获取用户信息
func (u *userService) GetUserInfos(userId int) (*model.UserInfos, error) {
	if userId <= 0 {
		return nil, nil
	}

	// 用户信息
	var userInfo *model.UserInfo
	err := dao.User.WherePri(userId).Scan(&userInfo)
	if err != nil {
		return nil, err
	}

	// 菜单资源
	menuResources, err := Resource.GetMenuByUserId(userId)
	if err != nil {
		return nil, err
	}
	menus := make([]model.ResourceInfo, 0, len(menuResources))
	for _, v := range menuResources {
		menus = append(menus, model.ResourceInfo{
			Id:       v.Id,
			ParentId: v.ParentId,
			Name:     v.Name,
			Alias:    v.Alias,
			Url:      v.Url,
			Icon:     v.Icon,
			Type:     v.Type,
			Sn:       v.Sn,
		})
	}

	userInfos := &model.UserInfos{
		User:  *userInfo,
		Menus: menus,
	}
	return userInfos, nil
}
