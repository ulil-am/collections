package genappid

import (
	"collections/helper/constant/tablename"
	dbStruct "collections/structs/db"
)

// Genappid - Logic Struct DB
type Genappid struct{}

var tblCounter = tablename.Counter

// GetInc - Get Increment
func (d *Genappid) GetInc(row dbStruct.Counter) (
	counter int,
	err error,
) {
	counter = 1
	return
}

// InsertServiceLog - InsertServiceLog
func (d *Genappid) InsertServiceLog(v interface{}) (err error) {

	return
}
