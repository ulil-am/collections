package cards

import (
	"collections/helper/constant"
	"collections/helper/constant/tablename"
	db "collections/models/db/mongo"
	structDb "collections/structs/db"

	"github.com/astaxie/beego"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Cards ...
type Cards struct{}

func init() {
	if constant.GOENV != constant.DEVCI {
		var d Cards
		d.Index()
	}
}

// Index ...
func (d *Cards) Index() (err error) {
	sess, coll, err := db.GetColl(tablename.Cards)
	defer sess.Close()
	if err != nil {
		return
	}

	index := mgo.Index{
		Key:        []string{"user_id", "card_number"},
		Unique:     true,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}

	err = coll.EnsureIndex(index)

	// index2 := mgo.Index{
	// 	Key:        []string{"user_id"},
	// 	Unique:     false,
	// 	DropDups:   false,
	// 	Background: false,
	// 	Sparse:     false,
	// }

	// err = coll.EnsureIndex(index2)
	return
}

// InsertCards ...
func (d *Cards) InsertCards(v interface{}) (err error) {
	sess, coll, err := db.GetColl(tablename.Cards)
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)
	return
}

// GetCardByCardNumber ...
func (d *Cards) GetCardByCardNumber(cardNumber int, u int) (rows []structDb.Cards, err error) {
	sess, coll, err := db.GetColl(tablename.Cards)
	defer sess.Close()
	if err != nil {
		return
	}
	err = coll.Find(bson.M{"user_id": u, "card_number": cardNumber}).All(&rows)
	beego.Debug("errr ======> ", err)

	return
}

// GetAllCards ...
func (d *Cards) GetAllCards(userID int) (rows []structDb.Cards, err error) {
	sess, coll, err := db.GetColl(tablename.Cards)
	defer sess.Close()
	if err != nil {
		return
	}
	err = coll.Find(bson.M{"user_id": userID}).All(&rows)

	return
}
