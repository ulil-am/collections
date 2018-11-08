package genappid

import (
	"medium/helper/constant"
)

func init() {
	initialize()
}

func initialize() {
	constant.LoadEnv()
}