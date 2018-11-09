package cards

import (
	"collections/helper/constant"
	iCards "collections/models/db/interfaces/cards"
	mgCards "collections/models/db/mongo/cards"
	stCards "collections/models/stub/cards"
)

var (
	logicName = "@Cards"
	// DBCards ...
	DBCards iCards.ICards
)

func init() {
	if constant.GOENV == constant.DEVCI {
		DBCards = new(stCards.Cards)
	} else {
		DBCards = new(mgCards.Cards)
	}
}
