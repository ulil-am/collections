package cards

import (
	"collections/helper/constant"
	iCards "collections/models/db/interfaces/cards"
	mgCards "collections/models/db/mongo/cards"
	stCards "collections/models/stub/cards"

	iUser "collections/models/db/interfaces/user"
	mgUser "collections/models/db/mongo/user"
)

var (
	logicName = "@Cards"
	// DBCards ...
	DBCards iCards.ICards
	// DBUser ...
	DBUser iUser.IUser
)

func init() {
	if constant.GOENV == constant.DEVCI {
		DBCards = new(stCards.Cards)
		// DBUser = new(stUser.User)
	} else {
		DBCards = new(mgCards.Cards)
		DBUser = new(mgUser.User)
	}
}
