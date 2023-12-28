package attraction

import (
	"context"
	"net/http"

	"inititaryplanner/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AttractionController struct {
	client *mongo.Client
}

func NewAttractionController(client *mongo.Client) *AttractionController {
	return &AttractionController{client}
}

func (ac AttractionController) GetAttraction(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	collection := ac.client.Database("mongo-golang").Collection("attractions")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attraction not found"})
		return
	}

	var attraction models.Attraction
	if err := result.Decode(&attraction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding JSON"})
		return
	}

	c.JSON(http.StatusOK, attraction)
}
