package event

import (
	"context"
	"encoding/json"
	"fmt"
	"gopractice/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ec EventController) GetEvents(w http.ResponseWriter, r *http.Request) {
	collection := ec.client.Database("mongo-golang").Collection("events")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error querying events:", err)
		return
	}
	defer cursor.Close(ctx)

	var events []models.Event
	for cursor.Next(ctx) {
		var e models.Event
		if err := cursor.Decode(&e); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error decoding JSON:", err)
			return
		}
		events = append(events, e)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error iterating through events:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(events); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
