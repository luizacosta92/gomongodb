package user

import "time"

type User struct {
	Name     string    `json:"name" bson:"name"`
	Whatsapp string    `json:"whatsapp" bson:"whatsapp"`
	Age      int       `json:"age" bson:"age"`
	Dpp      time.Time `json:"dpp" bson:"dpp"` //data provavel de parto
	City     string    `json:"city" bson:"city"` //cidade da gestante
}
