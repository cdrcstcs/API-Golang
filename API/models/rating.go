package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Type struct {
	Att bool
	Int bool
}
type Id struct {
	Value int
}
type Rating struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Type     Type               `bson:"type" json:"type"`
	UserId   Id                 `bson:"user_id" json:"user_id"`
	ObjectId Id                 `bson:"object_id" json:"object_id"`
	Score    int                `bson:"score" json:"score"`
}
