package user

import (
	"collections/helper/constant"
	"collections/helper/constant/tablename"
	db "collections/models/db/mongo"

	structDb "collections/structs/db"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// User ...
type User struct{}

func init() {
	if constant.GOENV != constant.DEVCI {
		var d User
		d.Index()
	}
}

// Index ...
func (d *User) Index() (err error) {
	sess, coll, err := db.GetColl(tablename.User)
	defer sess.Close()
	if err != nil {
		return
	}

	index := mgo.Index{
		Key:        []string{"user_name"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}

	err = coll.EnsureIndex(index)

	index2 := mgo.Index{
		Key:        []string{"user_id"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}

	err = coll.EnsureIndex(index2)
	return
}

// CreateUser ...
func (d *User) CreateUser(v interface{}) (err error) {
	sess, coll, err := db.GetColl(tablename.User)
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)
	return
}

// GetUserByUserName ...
func (d *User) GetUserByUserName(u string) (rows structDb.User, err error) {
	sess, coll, err := db.GetColl(tablename.User)
	defer sess.Close()
	if err != nil {
		return
	}
	err = coll.Find(bson.M{"user_name": u}).One(&rows)

	return
}
