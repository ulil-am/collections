package user

import (
	structDB "collections/structs/db"
)

// IUser - User Logic Interface
type IUser interface {
	CreateUser(v interface{}) (err error)
	GetUserByUserName(u string) (rows structDB.User, err error)
}
