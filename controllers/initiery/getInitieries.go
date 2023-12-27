package initiery

import (
	"context"
	"encoding/json"
	"fmt"
	"gopractice/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc InitieryController) GetInitieries(w http.ResponseWriter, r *http.Request) {
	collection := uc.client.Database("mongo-golang").Collection("initieries")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error querying initieries:", err)
		return
	}
	defer cursor.Close(ctx)

	var initieries []models.Initiery
	for cursor.Next(ctx) {
		var i models.Initiery
		if err := cursor.Decode(&i); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error decoding JSON:", err)
			return
		}
		initieries = append(initieries, i)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error iterating through initieries:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(initieries); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
