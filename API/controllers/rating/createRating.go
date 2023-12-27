package rating

import (
	"context"
	"fmt"
	"gopractice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func (rc RatingController) CreateRating(c *gin.Context) {
    var rating models.Rating
    if err := c.ShouldBindJSON(&rating); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error binding JSON: %s", err.Error())})
        return
    }

    collection := rc.client.Database("mongo-golang").Collection("ratings")
    ctx := context.TODO()

    _, err := collection.InsertOne(ctx, rating)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating rating: %s", err.Error())})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Rating created successfully", "rating": rating})
}
