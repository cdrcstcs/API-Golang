package rating

import (
	"context"
	"fmt"
	"net/http"

	"inititaryplanner/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RatingController struct {
	client *mongo.Client
}

func NewRatingController(client *mongo.Client) *RatingController {
	return &RatingController{client}
}

func (rc RatingController) GetRating(c *gin.Context) {
	var rating models.Rating

	// Extract rating ID from the request parameters
	ratingID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid rating ID: %s", err.Error())})
		return
	}

	collection := rc.client.Database("mongo-golang").Collection("ratings")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Rating not found: %s", result.Err().Error())})
		return
	}

	if err := result.Decode(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, rating)
}
