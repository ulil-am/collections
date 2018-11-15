package counter

import (
	structDb "collections/structs/db"
)

// ICounter - Counter Logic Interface
type ICounter interface {
	GetInc(row structDb.CounterUser) (
		counter int,
		err error,
	)
	InsertCounter(v interface{}) (err error)
}
