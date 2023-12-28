package attraction

import (
	"context"
	"fmt"
	"net/http"

	"inititaryplanner/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ac AttractionController) GetAttractions(c *gin.Context) {
	collection := ac.client.Database("mongo-golang").Collection("attractions")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error querying attractions: %s", err.Error())})
		return
	}
	defer cursor.Close(ctx)

	var attractions []models.Attraction
	for cursor.Next(ctx) {
		var attraction models.Attraction
		if err := cursor.Decode(&attraction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %s", err.Error())})
			return
		}
		attractions = append(attractions, attraction)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error iterating through attractions: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, attractions)
}
