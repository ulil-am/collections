package counter

import (
	"collections/helper/constant"
	"collections/helper/constant/tablename"
	db "collections/models/db/mongo"
	structDb "collections/structs/db"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Counter ...
type Counter struct{}

func init() {
	var d Counter
	d.Index()
}

// GetColl ...
func (d *Counter) GetColl() (sess *mgo.Session, coll *mgo.Collection, err error) {
	sess, err = db.Connect()
	if err != nil {
		beego.Warning("Error get collection Counter", err)
		return
	}

	coll = sess.DB(constant.GOAPP).C(tablename.User)

	return
}

// Index ...
func (d *Counter) Index() (err error) {
	sess, coll, err := db.GetColl(tablename.Counter)
	defer sess.Close()
	if err != nil {
		return
	}

	index := mgo.Index{
		Key:        []string{"counter"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}

	err = coll.EnsureIndex(index)
	return
}

// EXCLUSIVE DB OPERATIONS ===========================

// GetInc - Get Increment
func (d *Counter) GetInc(row structDb.CounterUser) (
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
		ReturnNew: true,
	}

	_, err = coll.Find(bson.M{"counter": row.Counter}).Apply(change, &row)
	counter = row.Counter
	beego.Debug("counter after inc", counter)

	return
}

// InsertCounter - InsertCounter
func (d *Counter) InsertCounter(v interface{}) (err error) {
	sess, coll, err := db.GetColl(tablename.Counter)
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)

	return
}
