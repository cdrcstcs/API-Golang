package event

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

type EventController struct {
	client *mongo.Client
}

func NewEventController(client *mongo.Client) *EventController {
	return &EventController{client}
}

func (ec EventController) GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := ec.client.Database("mongo-golang").Collection("events")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var event models.Event
	if err := result.Decode(&event); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(event); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
	}
}
