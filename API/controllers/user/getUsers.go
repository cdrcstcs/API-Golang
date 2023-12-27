package user

import (
	"context"
	"fmt"
	"gopractice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc UserController) GetUsers(c *gin.Context) {
	var users []models.User

	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error querying users: %s", err.Error())})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %s", err.Error())})
			return
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error iterating through users: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, users)
}
