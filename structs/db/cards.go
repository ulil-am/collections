package db

//Cards - Cards struct for DB
type Cards struct {
	CardNumber int    `bson:"card_number"`
	Name       string `bson:"name"`
	ExpiryDate string `bson:"expiry_date"`
	Company    string `bson:"company"`
	UserID     int    `bson:"user_id"`
}
