package attraction

import (
	"context"
	"encoding/json"
	"fmt"
	"gopractice/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ac AttractionController) GetAttractions(w http.ResponseWriter, r *http.Request) {
	collection := ac.client.Database("mongo-golang").Collection("attractions")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error querying attractions:", err)
		return
	}
	defer cursor.Close(ctx)

	var attractions []models.Attraction
	for cursor.Next(ctx) {
		var attraction models.Attraction
		if err := cursor.Decode(&attraction); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error decoding JSON:", err)
			return
		}
		attractions = append(attractions, attraction)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error iterating through attractions:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(attractions); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
