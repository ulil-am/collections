package genappid

import (
	"collections/helper/constant"
)

func init() {
	initialize()
}

func initialize() {
	constant.LoadEnv()
}