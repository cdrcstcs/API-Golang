package event

import (
	"context"
	"gopractice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventController struct {
	client *mongo.Client
}

func NewEventController(client *mongo.Client) *EventController {
	return &EventController{client}
}

func (ec EventController) GetEvent(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := ec.client.Database("mongo-golang").Collection("events")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var event models.Event
	if err := result.Decode(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	c.JSON(http.StatusOK, event)
}
