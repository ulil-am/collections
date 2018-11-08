package products

import (
	dbStruct "collections/structs/db"
)

// IProducts - Products Logic Interface
type IProducts interface {
	GetAllProduct() ([]dbStruct.Product, error)
	InsertProduct(interface{}) error
	GetProductByProductCodes([]string) (rows []dbStruct.Product, err error)
}
