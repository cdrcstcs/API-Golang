package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"

	"inititaryplanner/common/config"
)

func main() {
	flags := getFlags()
	config.InitGlobalConfig(flags.configPath)
	mc, err := InitializeMainController()
	if err != nil {
		os.Exit(1)
		return
	}

	r := gin.Default()
	//uc := user.NewUserController(client)
	//ac := attraction.NewAttractionController(client)
	//ec := event.NewEventController(client)
	//ic := itinerary.NewItineraryController(client)
	//rc := rating.NewRatingController(client)
	route(r, mc)

	// Start the Gin server
	err = r.Run("localhost:8100")
	if err != nil {
		os.Exit(1)
		return
	}
}

type flags struct {
	configPath string
}

func getFlags() flags {
	filePath := flag.String("config", "./config/local_config", "config path")
	flag.Parse()

	if filePath == nil || *filePath == "" {
		panic("empty config file path")
	}

	return flags{
		configPath: *filePath,
	}
}

func route(r *gin.Engine, m MainController) {
	// User routes
	//r.GET("/user/:id", uc.GetUser)
	//r.GET("/user", uc.GetUsers)
	//r.POST("/user", uc.CreateUser)
	r.POST("/user", m.CreateAttraction)

	// Attraction routes
	//r.GET("/attraction/:id", ac.GetAttraction)
	//r.GET("/attraction", ac.GetAttractions)
	//r.POST("/attraction", ac.CreateAttraction)

	// Event routes
	//r.GET("/event/:id", ec.GetEvent)
	//r.GET("/event", ec.GetEvents)
	//r.POST("/event", ec.CreateEvent)

	// Itinerary routes
	//r.GET("/itinerary/:id", ic.GetItinerary)
	//r.GET("/itinerary", ic.GetInitieries)
	//r.POST("/itinerary", ic.CreateItinerary)

	// Rating routes
	//r.GET("/rating/:id", rc.GetRating)
	//r.GET("/rating", rc.GetRatings)
	//r.POST("/rating", rc.CreateRating)

	// Add other routes as needed using r.GET or r.POST
}
