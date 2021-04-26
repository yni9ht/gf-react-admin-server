package service

import (
	"fmt"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestGetUserInfos(t *testing.T) {
	gtest.C(t, func(*gtest.T) {
		userInfos, _ := User.GetUserInfos(1)
		fmt.Printf("%v \n", userInfos)
	})
}
