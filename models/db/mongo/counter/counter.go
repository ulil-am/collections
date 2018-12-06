package counter

import (
	"collections/helper/constant/tablename"
	db "collections/models/db/mongo"
	structDb "collections/structs/db"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Counter ...
type Counter struct{}

func init() {}

// EXCLUSIVE DB OPERATIONS ===========================

// GetInc - Get Increment
func (d *Counter) GetInc(row structDb.CounterUser) (
	counter int,
	err error,
) {
	sess, coll, err := db.GetColl(tablename.CounterUser)
	defer sess.Close()
	if err != nil {
		return
	}

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"counter": 1}},
		ReturnNew: true,
	}

	_, err = coll.Find(bson.M{"_id": row.UserCounter}).Apply(change, &row)
	counter = row.Counter
	// beego.Debug("info after inc", info)
	beego.Debug("counter after inc", counter)
	return
}

// InsertCounter - InsertCounter
func (d *Counter) InsertCounter(v interface{}) (err error) {
	sess, coll, err := db.GetColl(tablename.CounterUser)
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)

	return
}
