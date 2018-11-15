package db

// User ...
type User struct {
	UserName string `bson:"user_name"`
	UserID   int    `bson:"user_id"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
}
