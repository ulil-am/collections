package cards

import structDb "collections/structs/db"

// ICards ...
type ICards interface {
	InsertCards(interface{}) error
	GetCardByCardNumber(string) (rows []structDb.Cards, err error)
}
