package initiery

import (
	"context"
	"fmt"
	"net/http"

	"inititaryplanner/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ic InitieryController) GetInitiery(c *gin.Context) {
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Initiery not found"})
		return
	}

	var initiery models.Initiery
	if err := result.Decode(&initiery); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, initiery)
}
