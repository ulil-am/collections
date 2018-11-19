package cards

import structDb "collections/structs/db"

// ICards ...
type ICards interface {
	InsertCards(interface{}) error
	GetCardByCardNumber(int) ([]structDb.Cards, error)
	GetAllCards(int) ([]structDb.Cards, error)
}
