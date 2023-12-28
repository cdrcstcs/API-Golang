package dal

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"inititaryplanner/common/custom_errs"
	"inititaryplanner/common/utils"
	"inititaryplanner/constant"
	"inititaryplanner/dal/db"
	"inititaryplanner/dal/inf"
	"inititaryplanner/models"
)

func GetAttractionDal(mainDB *db.MainMongoDB) inf.AttractionDal {
	return &attractionDal{
		mainDB: (*mongo.Database)(mainDB),
	}
}

type attractionDal struct {
	mainDB *mongo.Database
}

func (a *attractionDal) CreateAttraction(ctx context.Context, attraction *models.Attraction) (*models.Attraction, error) {
	if !utils.IsEmpty(attraction.Id) {
		// TODO logging here
		return nil, custom_errs.DBErrCreateWithID
	}
	collection := a.mainDB.Collection(constant.AttractionTable)
	// Insert newAttraction into MongoDB
	result, err := collection.InsertOne(ctx, attraction)
	if err != nil {
		// TODO logging here
		return nil, err
	}

	// Extract inserted ID
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		// TODO logging here
		return nil, custom_errs.DBErrIDConversion
	}

	attraction.Id = insertedID.String()
	return attraction, nil
}
