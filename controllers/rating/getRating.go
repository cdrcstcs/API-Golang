package rating

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

type RatingController struct {
	client *mongo.Client
}

func NewRatingController(client *mongo.Client) *RatingController {
	return &RatingController{client}
}

func (rc RatingController) GetRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := rc.client.Database("mongo-golang").Collection("ratings")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var rating models.Rating
	if err := result.Decode(&rating); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(rating); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
	}
}
