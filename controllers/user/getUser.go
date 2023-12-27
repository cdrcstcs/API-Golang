package user

import (
	"context"
	"encoding/json"
	"fmt"
	"gopractice/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var u models.User
	if err := result.Decode(&u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
	}
}
