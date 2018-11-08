package cards

import (
	"collections/helper/constant"
	"collections/helper/constant/tablename"
	db "collections/models/db/mongo"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
)

// Cards ...
type Cards struct{}

func init() {
	var d Cards
	d.Index()
}

// GetColl ...
func (d *Cards) GetColl() (sess *mgo.Session, coll *mgo.Collection, err error) {
	sess, err = db.Connect()
	if err != nil {
		beego.Warning("Error get collection Cards", err)
		return
	}

	coll = sess.DB(constant.GOAPP).C(tablename.Cards)

	return
}

// Index ...
func (d *Cards) Index() (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	index := mgo.Index{
		Key:        []string{"card_number"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}

	err = coll.EnsureIndex(index)
	return
}

// InsertCards ...
func (d *Cards) InsertCards(v interface{}) (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)
	return
}
