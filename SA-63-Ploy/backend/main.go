package main

import (
	"context"
	"log"

	"github.com/Sujitnapa21/app/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Sujitnapa21/app/ent"
)

// Customers  defines the struct for the customers
type Customers struct {
	Customer []Customer
}

// Customer  defines the struct for the customer
type Customer struct {
	Email    string
	Username string
}

// Rooms  defines the struct for the rooms
type Rooms struct {
	Room []Room
}

// Room  defines the struct for the  room
type Room struct {
	Roomtype string
}

// Adults  defines the struct for the ddults
type Adults struct {
	Adult []Adult
}

// Adult  defines the struct for the adult
type Adult struct {
	Amount int
}

// Kids  defines the struct for the kids
type Kids struct {
	Kid []Kid
}

// Kid  defines the struct for the kid
type Kid struct {
	Amount int
}

// Roomamounts  defines the struct for the roomamounts
type Roomamounts struct {
	Roomamount []Roomamount
}

// Roomamount  defines the struct for the roomamount
type Roomamount struct {
	Amount int
}

// @title SUT SA Example API Patient
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:Books.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewBooksController(v1, client)
	controllers.NewCustomerController(v1, client)
	controllers.NewRoomController(v1, client)
	controllers.NewAdultController(v1, client)
	controllers.NewKidController(v1, client)
	controllers.NewRoomamountController(v1, client)

	// Set Customers Data
	customers := Customers{
		Customer: []Customer{
			Customer{"me@gmail.com", "Me123"},
		},
	}

	for _, cus := range customers.Customer {
		client.Customer.
			Create().
			SetEmail(cus.Email).
			SetUsername(cus.Username).
			Save(context.Background())
	}

	// Set Rooms Data
	rooms := Rooms{
		Room: []Room{
			Room{"Standard"},
			Room{"Superior"},
			Room{"Duluxe"},
		},
	}

	for _, r := range rooms.Room {
		client.Room.
			Create().
			SetRoomtype(r.Roomtype).
			Save(context.Background())
	}

	// Set Adults Data
	adults := Adults{
		Adult: []Adult{
			Adult{1},
			Adult{2},
			Adult{3},
		},
	}

	for _, a := range adults.Adult {
		client.Adult.
			Create().
			SetAmount(a.Amount).
			Save(context.Background())
	}

	// Set Kids Data
	kids := Kids{
		Kid: []Kid{
			Kid{1},
			Kid{2},
			Kid{3},
		},
	}

	for _, k := range kids.Kid {
		client.Kid.
			Create().
			SetAmount(k.Amount).
			Save(context.Background())
	}

	// Set Roomamount Data
	roomamounts := Roomamounts{
		Roomamount: []Roomamount{
			Roomamount{1},
			Roomamount{2},
			Roomamount{3},
		},
	}

	for _, ra := range roomamounts.Roomamount {
		client.Roomamount.
			Create().
			SetAmount(ra.Amount).
			Save(context.Background())
	}

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
