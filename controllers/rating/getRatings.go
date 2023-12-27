package rating

import (
	"context"
	"encoding/json"
	"fmt"
	"gopractice/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (rc RatingController) GetRatings(w http.ResponseWriter, r *http.Request) {
	collection := rc.client.Database("mongo-golang").Collection("ratings")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error querying ratings:", err)
		return
	}
	defer cursor.Close(ctx)

	var ratings []models.Rating
	for cursor.Next(ctx) {
		var rating models.Rating
		if err := cursor.Decode(&rating); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error decoding JSON:", err)
			return
		}
		ratings = append(ratings, rating)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error iterating through ratings:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ratings); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
