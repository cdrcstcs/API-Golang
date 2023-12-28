package models

type Coordinate struct {
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`
}

type Tag struct {
	Id    string `bson:"_id" json:"id"`
	Value string `bson:"value" json:"value"` // value should be unique, however tag value should not be passes around
}

type Attraction struct {
	Id         string     `bson:"_id" json:"id"`
	Name       string     `bson:"name" json:"name"`
	Address    string     `bson:"address" json:"address"`
	Coordinate Coordinate `bson:"coordinate" json:"coordinate"`
	TagIDs     []string   `bson:"tag_ids" json:"tag_ids"`
}
