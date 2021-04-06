package main

import (
	_ "gf-react-admin-server/boot"
	_ "gf-react-admin-server/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
