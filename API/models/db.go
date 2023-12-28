package models

import "time"

type Coordinate struct {
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`
}

type Tag struct {
	Id    string `bson:"_id,omitempty" json:"id"`
	Value string `bson:"value" json:"value"` // value should be unique, however tag value should not be passes around
}

type Attraction struct {
	Id         string     `bson:"_id,omitempty" json:"id"`
	Name       string     `bson:"name" json:"name"`
	Address    string     `bson:"address" json:"address"`
	Coordinate Coordinate `bson:"coordinate" json:"coordinate"`
	TagIDs     []string   `bson:"tag_ids" json:"tag_ids"`
}

type User struct {
	Id            string `bson:"_id,omitempty" json:"id"`
	Name          string `bson:"name" json:"name"`
	Password      string `bson:"password" json:"password"`
	Email         string `bson:"email" json:"email"`
	EmailPassword string `bson:"email_password" json:"email_password"`
}

type Type int64

func (t Type) String() string {
	switch t {
	case TypeAttraction:
		return "ATTRACTION"
	case TypeItinerary:
		return "ITINERARY"
	default:
		return "<UNSET>"
	}
}

const (
	TypeAttraction Type = 1
	TypeItinerary  Type = 2
)

type Id struct {
	Value int
}
type Rating struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Type     Type   `bson:"type" json:"type"`
	UserId   string `bson:"user_id" json:"user_id"`
	ObjectId string `bson:"object_id" json:"object_id"`
	Score    int    `bson:"score" json:"score"`
}

type Event struct {
	Id          string     `bson:"_id,omitempty" json:"id"`
	ItineraryId Id         `bson:"itinerary_id" json:"itinerary_id"`
	StartTime   time.Time  `bson:"start_time" json:"start_time"`
	EndTime     time.Time  `bson:"end_time" json:"end_time"`
	Attraction  Attraction `bson:"attraction" json:"attraction"`
	Description string     `bson:"description" json:"description"`
}

type Itinerary struct {
	Id        string    `bson:"_id,omitempty" json:"id"`
	CopiedId  Id        `bson:"copied_id" json:"copied_id"`
	UserId    Id        `bson:"user_id" json:"user_id"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
	Events    []Event   `bson:"events" json:"events"`
	Rating    Rating    `bson:"rating" json:"rating"`
}
