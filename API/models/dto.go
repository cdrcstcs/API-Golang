package models

type AttractionDTO struct {
	Id         string     `bson:"_id" json:"id"`
	Name       string     `bson:"name" json:"name"`
	Address    string     `bson:"address" json:"address"`
	Coordinate Coordinate `bson:"coordinate" json:"coordinate"`
	Tags       []Tag      `bson:"tags" json:"tags"`
}

type CreateAttractionReq struct {
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	Coordinate Coordinate `json:"coordinate"`
	TagIDs     []string   `json:"tag_ids"`
}

type CreateAttractionResp struct {
	Attraction *AttractionDTO `json:"attraction"`
}
