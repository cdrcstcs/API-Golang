package initiery

import (
	"context"
	"fmt"
	"gopractice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ic InitieryController) CreateInitiery(c *gin.Context) {
	var newInitiery models.Initiery

	// Bind JSON request body to newInitiery
	if err := c.ShouldBindJSON(&newInitiery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error decoding request body: %s", err.Error())})
		return
	}

	collection := ic.client.Database("mongo-golang").Collection("initieries")
	ctx := context.TODO()

	// Insert newInitiery into MongoDB
	result, err := collection.InsertOne(ctx, newInitiery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error inserting initiery: %s", err.Error())})
		return
	}

	// Extract the ID generated by MongoDB after insertion
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error extracting inserted ID"})
		return
	}

	newInitiery.Id = insertedID

	// Respond with the created initiery
	c.JSON(http.StatusCreated, newInitiery)
}
