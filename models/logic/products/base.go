package products

import (
	"collections/helper/constant"
	iProducts "collections/models/db/interfaces/products"
	mgProducts "collections/models/db/mongo/products"
	stProducts "collections/models/stub/products"
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
