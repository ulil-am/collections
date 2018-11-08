package genappid

import (
	"medium/helper/constant/tablename"
	db "medium/models/db/mongo"
	dbStruct "medium/structs/db"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Genappid - Logic Struct DB
type Genappid struct{}

func init() {
	var d Genappid
	d.Index()
}

// Index - Create Index
func (d *Genappid) Index() (err error) {
	sess, coll, err := db.GetColl(tablename.Counter)
	defer sess.Close()
	if err != nil {
		return
	}

	index := mgo.Index{
		Key:        []string{"branch_code", "month"},
		Unique:     true,  // Prevent two documents from having the same index key
		DropDups:   false, // Drop documents with the same index key as a previously indexed one
		Background: false, // Build index in background and return immediately
		Sparse:     false, // Only index documents containing the Key fields
	}

	err = coll.EnsureIndex(index)

	return
}

// GetInc - Get Increment
func (d *Genappid) GetInc(row dbStruct.Counter) (
	counter int,
	err error,
) {
	sess, coll, err := db.GetColl(tablename.Counter)
	defer sess.Close()
	if err != nil {
		return
	}

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counter": 1}},
		ReturnNew: false,
	}

	_, err = coll.Find(bson.M{"branch_code": row.BranchCode, "month": row.Month}).Apply(change, &row)
	counter = row.Counter

	return
}

// InsertServiceLog - InsertServiceLog
func (d *Genappid) InsertServiceLog(v interface{}) (err error) {
	sess, coll, err := db.GetColl(tablename.Counter)
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)

	return
}
