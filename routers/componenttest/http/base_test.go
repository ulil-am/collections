package http

import (
	"collections/routers/componenttest"
	"strconv"

	"github.com/astaxie/beego"
)

var (
	host string = "http://127.0.0.1:" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)
)

func init() {
	componenttest.HTTPInit()
	// componenttest.DBinit()
}
