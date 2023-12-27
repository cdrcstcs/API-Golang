package rating

import (
	"context"
	"fmt"
	"gopractice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (rc RatingController) GetRatings(c *gin.Context) {
    collection := rc.client.Database("mongo-golang").Collection("ratings")
    ctx := context.TODO()

    cursor, err := collection.Find(ctx, primitive.D{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error querying ratings: %s", err.Error())})
        return
    }
    defer cursor.Close(ctx)

    var ratings []models.Rating
    for cursor.Next(ctx) {
        var rating models.Rating
        if err := cursor.Decode(&rating); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %s", err.Error())})
            return
        }
        ratings = append(ratings, rating)
    }

    if err := cursor.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error iterating through ratings: %s", err.Error())})
        return
    }

    c.JSON(http.StatusOK, ratings)
}
