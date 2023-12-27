package main

import (
	"context"
	"gopractice/controllers/attraction"
	"gopractice/controllers/event"
	"gopractice/controllers/initiery"
	"gopractice/controllers/rating"
	"gopractice/controllers/user"

	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := mux.NewRouter()
	client := getClient()
	uc := user.NewUserController(client)
	ac := attraction.NewAttractionController(client)
	ec := event.NewEventController(client)
	ic := initiery.NewInitieryController(client)
	rc := rating.NewRatingController(client)
	r.HandleFunc("/user/{id}", uc.GetUser).Methods("GET")
	r.HandleFunc("/user", uc.GetUsers).Methods("GET")
	r.HandleFunc("/user", uc.CreateUser).Methods("POST")
	r.HandleFunc("/attraction/{id}", ac.GetAttraction).Methods("GET")
	r.HandleFunc("/attraction", ac.GetAttractions).Methods("GET")
	r.HandleFunc("/attraction", ac.CreateAttraction).Methods("POST")
	r.HandleFunc("/event/{id}", ec.GetEvent).Methods("GET")
	r.HandleFunc("/event", ec.GetEvents).Methods("GET")
	r.HandleFunc("/event", ec.CreateEvent).Methods("POST")
	r.HandleFunc("/initiery/{id}", ic.GetInitiery).Methods("GET")
	r.HandleFunc("/initiery", ic.GetInitieries).Methods("GET")
	r.HandleFunc("/initiery", ic.CreateInitiery).Methods("POST")
	r.HandleFunc("/rating/{id}", rc.GetRating).Methods("GET")
	r.HandleFunc("/rating", rc.GetRatings).Methods("GET")
	r.HandleFunc("/rating", rc.CreateRating).Methods("POST")
	// Add other routes as needed using r.HandleFunc

	http.Handle("/", r)
	http.ListenAndServe("localhost:8100", nil)
}

func getClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return client
}
