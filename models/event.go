package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	ItineraryId Id                 `bson:"itinerary_id" json:"itinerary_id"`
	StartTime   Date               `bson:"start_time" json:"start_time"`
	EndTime     Date               `bson:"end_time" json:"end_time"`
	Attraction  Attraction         `bson:"attraction" json:"attraction"`
	Description string             `bson:"description" json:"description"`
}
