package products

import (
	"medium/helper/constant/tablename"
	dbStruct "medium/structs/db"
	"encoding/json"
	"errors"
)

// Products - Logic Struct DB
type Products struct{}

var tblProduct = tablename.Product

// GetAllProduct - GetAllProduct GetAll
func (d *Products) GetAllProduct() (rows []dbStruct.Product, err error) {
	// row := dbStruct.Product{
	// 	JobID:   "job1",
	// 	Req:     "req",
	// 	Res:     "res",
	// 	Errcode: "errcode",
	// 	Type:    "http",
	// }
	// rows = append(rows, row)
	return
}

// GetProductByProductCodes - Get product by product codes
func (d *Products) GetProductByProductCodes(productCodes []string) (rows []dbStruct.Product, err error) {
	for _, val := range productCodes {
		if val == "productError" {
			err = errors.New("fail to get product")
			return
		}
		rows = append(rows, dbStruct.Product{
			CampaignCode: val,
		})
	}
	return
}

// InsertProduct - InsertProduct Insert
func (d *Products) InsertProduct(v interface{}) (err error) {
	var row dbStruct.Product
	interfaceProduct, err := json.Marshal(v)
	if err != nil {
		return
	}

	err = json.Unmarshal(interfaceProduct, &row)
	if err != nil {
		return
	}

	if row.CampaignCode == "productError" {
		err = errors.New("fail to get product")
		return
	}

	return
}
