package event

import (
	"context"
	"fmt"
	"gopractice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ec EventController) GetEvents(c *gin.Context) {
	collection := ec.client.Database("mongo-golang").Collection("events")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error querying events: %s", err.Error())})
		return
	}
	defer cursor.Close(ctx)

	var events []models.Event
	for cursor.Next(ctx) {
		var e models.Event
		if err := cursor.Decode(&e); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %s", err.Error())})
			return
		}
		events = append(events, e)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error iterating through events: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, events)
}
