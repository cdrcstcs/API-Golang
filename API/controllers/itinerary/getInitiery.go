package itinerary

import (
	"context"
	"fmt"
	"net/http"

	"inititaryplanner/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ic ItineraryController) GetItinerary(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := ic.client.Database("mongo-golang").Collection("initieries")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Itinerary not found"})
		return
	}

	var itinerary models.Itinerary
	if err := result.Decode(&itinerary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, itinerary)
}
