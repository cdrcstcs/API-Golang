package main

import (
	"context"

	"inititaryplanner/controllers/itinerary"

	"inititaryplanner/controllers/attraction"
	"inititaryplanner/controllers/event"
	"inititaryplanner/controllers/rating"
	"inititaryplanner/controllers/user"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	client := getClient()
	uc := user.NewUserController(client)
	ac := attraction.NewAttractionController(client)
	ec := event.NewEventController(client)
	ic := itinerary.NewItineraryController(client)
	rc := rating.NewRatingController(client)

	// User routes
	r.GET("/user/:id", uc.GetUser)
	r.GET("/user", uc.GetUsers)
	r.POST("/user", uc.CreateUser)

	// Attraction routes
	r.GET("/attraction/:id", ac.GetAttraction)
	r.GET("/attraction", ac.GetAttractions)
	r.POST("/attraction", ac.CreateAttraction)

	// Event routes
	r.GET("/event/:id", ec.GetEvent)
	r.GET("/event", ec.GetEvents)
	r.POST("/event", ec.CreateEvent)

	// Itinerary routes
	r.GET("/itinerary/:id", ic.GetItinerary)
	r.GET("/itinerary", ic.GetInitieries)
	r.POST("/itinerary", ic.CreateItinerary)

	// Rating routes
	r.GET("/rating/:id", rc.GetRating)
	r.GET("/rating", rc.GetRatings)
	r.POST("/rating", rc.CreateRating)

	// Add other routes as needed using r.GET or r.POST

	// Start the Gin server
	r.Run("localhost:8100")
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
