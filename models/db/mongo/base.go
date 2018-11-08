package mongo

import (
	"collections/helper"
	"collections/helper/constant"

	"github.com/globalsign/mgo"
)

// Connect to connect Mongo DB
// this is use for open connection to mongodb
func Connect() (sess *mgo.Session, err error) {
	sess, err = mgo.Dial(constant.CREDMONGO)
	if err != nil {
		helper.CheckErr("error connect mongo DB", err)
		return
	}

	mgo.SetDebug(true)

	sess.SetMode(mgo.Monotonic, true)

	return
}

// GetColl - Get Collection
func GetColl(collName string) (sess *mgo.Session, coll *mgo.Collection, err error) {
	sess, err = Connect()
	if err != nil {
		helper.CheckErr("Failed get collection "+collName, err)
		return
	}

	coll = sess.DB(constant.GOAPP).C(collName)

	return
}
