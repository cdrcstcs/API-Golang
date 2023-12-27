package user

import (
	"context"
	"encoding/json"
	"fmt"
	"gopractice/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error querying users:", err)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var u models.User
		if err := cursor.Decode(&u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error decoding JSON:", err)
			return
		}
		users = append(users, u)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error iterating through users:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
