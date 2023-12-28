package main

import (
	"context"

	"inititaryplanner/controllers/attraction"
	"inititaryplanner/controllers/event"
	"inititaryplanner/controllers/initiery"
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
	ic := initiery.NewInitieryController(client)
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

	// Initiery routes
	r.GET("/initiery/:id", ic.GetInitiery)
	r.GET("/initiery", ic.GetInitieries)
	r.POST("/initiery", ic.CreateInitiery)

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
