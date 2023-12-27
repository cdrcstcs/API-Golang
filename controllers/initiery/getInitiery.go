package initiery

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

type InitieryController struct {
	client *mongo.Client
}

func NewInitieryController(client *mongo.Client) *InitieryController {
	return &InitieryController{client}
}

func (ic InitieryController) GetInitiery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	collection := ic.client.Database("mongo-golang").Collection("initieries")
	ctx := context.TODO()

	result := collection.FindOne(ctx, primitive.M{"_id": oid})
	if result.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var initiery models.Initiery
	if err := result.Decode(&initiery); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(initiery); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error encoding JSON:", err)
	}
}
