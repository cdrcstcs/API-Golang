package attraction

import (
	"context"
	"fmt"
	"gopractice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ac AttractionController) CreateAttraction(c *gin.Context) {
	var newAttraction models.Attraction

	// Bind JSON request body to newAttraction
	if err := c.ShouldBindJSON(&newAttraction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}

	collection := ac.client.Database("mongo-golang").Collection("attractions")
	ctx := context.TODO()

	// Insert newAttraction into MongoDB
	result, err := collection.InsertOne(ctx, newAttraction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error inserting attraction: %s", err.Error())})
		return
	}

	// Extract inserted ID
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error extracting inserted ID"})
		return
	}

	newAttraction.Id = insertedID

	// Respond with the created attraction
	c.JSON(http.StatusCreated, newAttraction)
}
