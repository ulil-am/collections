package cards

import (
	"collections/helper/constant/tablename"
	structDb "collections/structs/db"
	"encoding/json"
	"errors"
)

// Cards ...
type Cards struct{}

var tblProduct = tablename.Cards

// InsertCards - InsertCards Insert
func (d *Cards) InsertCards(v interface{}) (err error) {
	var rows structDb.Cards
	interfaceCards, err := json.Marshal(v)
	if err != nil {
		return
	}

	err = json.Unmarshal(interfaceCards, &rows)
	if err != nil {
		return
	}

	if rows.CardNumber == 0 {
		err = errors.New("fail to insert cards")
		return
	}

	return
}

// GetCardByCardNumber ...
func (d *Cards) GetCardByCardNumber(cardNumber int) (rows []structDb.Cards, err error) {
	var row structDb.Cards

	if cardNumber == 0 {
		err = errors.New("fail to get cards")
		return
	}

	row.CardNumber = cardNumber
	row.Company = "TN"
	row.ExpiryDate = "2018/11/11"
	row.Name = "Akira Kaido"

	rows = append(rows, row)
	return
}
