package products

import (
	"medium/helper/constant"
	iProducts "medium/models/db/interfaces/products"
	mgProducts "medium/models/db/mongo/products"
	stProducts "medium/models/stub/products"
)

var (
	logicName = "@products"
	//DBProducts ...
	DBProducts iProducts.IProducts
)

func init() {
	if constant.GOENV == constant.DEVCI {
		DBProducts = new(stProducts.Products)
	} else {
		DBProducts = new(mgProducts.Products)
	}
}
