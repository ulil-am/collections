package user

// IUser - User Logic Interface
type IUser interface {
	CreateUser(v interface{}) (err error)
}
