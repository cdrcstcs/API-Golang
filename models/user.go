package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Password      string             `bson:"password" json:"password"`
	Email         string             `bson:"email" json:"email"`
	EmailPassword string             `bson:"email_password" json:"email_password"`
}
