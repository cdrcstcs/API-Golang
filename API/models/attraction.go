package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coordinate struct {
	X int
	Y int
}
type Tags []Tag
type Tag struct {
	Value string
}
type Attraction struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string
	Address    string
	Coordinate Coordinate
	Tags       Tags
}
