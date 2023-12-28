package initiery

import (
	"context"
	"fmt"
	"net/http"

	"inititaryplanner/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InitieryController struct {
	client *mongo.Client
}

func NewInitieryController(client *mongo.Client) *InitieryController {
	return &InitieryController{client}
}

// GetInitieries handles the GET request for retrieving initieries
func (ic InitieryController) GetInitieries(c *gin.Context) {
	collection := ic.client.Database("mongo-golang").Collection("initieries")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying initieries"})
		fmt.Println("Error querying initieries:", err)
		return
	}
	defer cursor.Close(ctx)

	var initieries []models.Initiery
	for cursor.Next(ctx) {
		var i models.Initiery
		if err := cursor.Decode(&i); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
			fmt.Println("Error decoding JSON:", err)
			return
		}
		initieries = append(initieries, i)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating through initieries"})
		fmt.Println("Error iterating through initieries:", err)
		return
	}

	c.JSON(http.StatusOK, initieries)
}
