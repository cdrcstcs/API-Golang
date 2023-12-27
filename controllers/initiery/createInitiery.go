package initiery

import (
	"context"
	"encoding/json"
	"fmt"
	"gopractice/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ic InitieryController) CreateInitiery(w http.ResponseWriter, r *http.Request) {
	var newInitiery models.Initiery
	err := json.NewDecoder(r.Body).Decode(&newInitiery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error decoding request body:", err)
		return
	}

	collection := ic.client.Database("mongo-golang").Collection("initieries")
	ctx := context.TODO()

	result, err := collection.InsertOne(ctx, newInitiery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error inserting initiery:", err)
		return
	}

	// Extract the ID generated by MongoDB after insertion
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error extracting inserted ID")
		return
	}

	newInitiery.Id = insertedID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newInitiery); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON: ", err)
		return
	}
}
