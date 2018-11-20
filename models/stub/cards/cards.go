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
func (d *Cards) GetCardByCardNumber(cardNumber int, userID int) (rows []structDb.Cards, err error) {
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

// GetAllCards ...
func (d *Cards) GetAllCards(userID int) (rows []structDb.Cards, err error) {
	var row structDb.Cards

	if userID == 0 {
		err = errors.New("user is not exists")
	}

	cardNumber := []int{21313111, 4123111}
	company := []string{"TN", "TN"}
	expiryDate := []string{"2018/11/11", "2018/11/11"}
	name := []string{"Akira Kaido", "Akira Kaido"}

	for i := 0; i < 2; i++ {
		row = structDb.Cards{
			CardNumber: cardNumber[i],
			Company:    company[i],
			ExpiryDate: expiryDate[i],
			Name:       name[i],
		}
	}

	rows = append(rows, row)
	return
}
