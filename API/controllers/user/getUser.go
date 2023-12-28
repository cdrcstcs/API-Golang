package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"inititaryplanner/models"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}
func (uc UserController) GetUser(c *gin.Context) {
	var user models.User
	userID := c.Param("id") // Assuming you get the user ID from the request parameters
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid user ID: %s", err.Error())})
		return
	}

	collection := uc.client.Database("mongo-golang").Collection("users") // Assuming your user collection is named "users"
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", result.Err().Error())})
		return
	}

	if err := result.Decode(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding JSON: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, user)
}
