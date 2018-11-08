package genappid

import (
	dbStruct "medium/structs/db"
)

// IGenappid - Genappid Logic Interface
type IGenappid interface {
	GetInc(row dbStruct.Counter) (
		counter int,
		err error,
	)
	InsertServiceLog(v interface{}) (err error)
}
