package service

import (
	"gf-vue3-admin-server/app/dao"
	"gf-vue3-admin-server/app/model"
)

type userService struct {
}

var (
	User = &userService{}
)

func (u *userService) Login(loginReq *model.UserLoginReq) (user *model.UserLoginRes, err error) {
	var queryParam = map[string]string{
		"account":  loginReq.Account,
		"password": loginReq.Password,
	}
	err = dao.User.Where(queryParam).Scan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}