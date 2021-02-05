package main

import (
	_ "gf-vue3-admin-server/boot"
	_ "gf-vue3-admin-server/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
