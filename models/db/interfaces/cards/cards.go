package cards

import structDb "collections/structs/db"

// ICards ...
type ICards interface {
	InsertCards(interface{}) error
	GetCardByCardNumber(int) (rows []structDb.Cards, err error)
}
