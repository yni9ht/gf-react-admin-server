package service

import (
	"fmt"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestGetUserMenu(t *testing.T) {
	gtest.C(t, func(*gtest.T) {
		menus, _ := Resource.GetMenuByUserId(1)
		fmt.Printf("%v \n", menus)
	})
}

func TestGetUserMenuTree(t *testing.T) {
	gtest.C(t, func(*gtest.T) {
		treeNode, _ := Resource.GetMenuTreeByUserId(1)
		fmt.Print(treeNode)
	})
}
