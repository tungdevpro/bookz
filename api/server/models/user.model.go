package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Name    string             `bson:"name" json:"name"`
	Phone   string             `bson:"phone" json:"phone"`
	Address string             `bson:"address" json:"address"`
}
