package products

import (
	"collections/helper/constant"
)

func init() {
	initialize()
}

func initialize() {
	constant.LoadEnv()
}