package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

type Initiery struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	CopiedId  Id                 `bson:"copied_id" json:"copied_id"`
	UserId    Id                 `bson:"user_id" json:"user_id"`
	StartTime Date               `bson:"start_time" json:"start_time"`
	EndTime   Date               `bson:"end_time" json:"end_time"`
	Events    []Event            `bson:"events" json:"events"`
	Rating    Rating             `bson:"rating" json:"rating"`
}
