package products

import (
	"medium/helper"
	"medium/helper/constant"
	"medium/helper/constant/tablename"
	db "medium/models/db/mongo"
	dbStruct "medium/structs/db"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Products - Logic Struct DB
type Products struct{}

func init() {
	var d Products
	d.Index()
}

// GetColl - Get Collection product
func (d *Products) GetColl() (sess *mgo.Session, coll *mgo.Collection, err error) {
	sess, err = db.Connect()
	if err != nil {
		helper.CheckErr("Failed get collection product", err)
		return
	}

	coll = sess.DB(constant.GOAPP).C(tablename.Product)

	return
}

// Index - Create Index
func (d *Products) Index() (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	index := mgo.Index{
		Key:        []string{"campaign_code"},
		Unique:     true,  // Prevent two documents from having the same index key
		DropDups:   false, // Drop documents with the same index key as a previously indexed one
		Background: false, // Build index in background and return immediately
		Sparse:     false, // Only index documents containing the Key fields
	}

	err = coll.EnsureIndex(index)

	return
}

// GetAllProduct - GetAllProduct GetAll
func (d *Products) GetAllProduct() (rows []dbStruct.Product, err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Find(bson.M{}).All(&rows)

	return
}

// GetProductByProductCodes - Get product by product codes
func (d *Products) GetProductByProductCodes(productCodes []string) (rows []dbStruct.Product, err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Find(bson.M{"campaign_code": bson.M{"$in": productCodes}}).All(&rows)

	return
}

// GetOneByProductCode...
func (d *Products) GetOneByProductCode() (row dbStruct.Product, err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Find(bson.M{"campaign_code": row.CampaignCode}).One(&row)

	return
}

// UpdateByProductCode...
func (d *Products) UpdateByProductCode() (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	selector := bson.M{"campaign_code": "xxxxxx"}
	update := bson.M{
		"$set": bson.M{
			"res": "yyyy",
		},
	}

	err = coll.Update(selector, update)

	return
}

// InsertProduct - InsertProduct
func (d *Products) InsertProduct(v interface{}) (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)

	return
}

// DeleteByProductCode - DeleteByProductCode
func (d *Products) DeleteByProductCode() (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}
	selector := bson.M{"job_id": "xxxxxx"}
	err = coll.Remove(selector)

	return
}
