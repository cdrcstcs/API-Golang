package attraction

import (
	"context"
	"encoding/json"
	"fmt"
	"gopractice/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ac AttractionController) CreateAttraction(w http.ResponseWriter, r *http.Request) {
	var newAttraction models.Attraction
	err := json.NewDecoder(r.Body).Decode(&newAttraction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error decoding request body: ", err)
		return
	}
	collection := ac.client.Database("mongo-golang").Collection("attractions")
	ctx := context.TODO()
	result, err := collection.InsertOne(ctx, newAttraction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error inserting attraction: ", err)
		return
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error extracting inserted ID")
		return
	}
	newAttraction.Id = insertedID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newAttraction); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON: ", err)
		return
	}
}
