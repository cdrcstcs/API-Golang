package attraction

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

type AttractionController struct {
	client *mongo.Client
}

func NewAttractionController(client *mongo.Client) *AttractionController {
	return &AttractionController{client}
}

func (ac AttractionController) GetAttraction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := ac.client.Database("mongo-golang").Collection("attractions")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var attraction models.Attraction
	if err := result.Decode(&attraction); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(attraction); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
	}
}
