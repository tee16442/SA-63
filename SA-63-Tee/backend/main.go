package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tee16/app/controllers"
	"github.com/tee16/app/ent"
)

// Problemtypes  defines the struct for the problemtypes
type Problemtypes struct {
	Problemtype []Problemtype
}

// Problemtype  defines the struct for the  problemtype
type Problemtype struct {
	problemtype string
}

// Rooms  defines the struct for the Rooms
type Rooms struct {
	Room []Room
}

// Room  defines the struct for the  Room
type Room struct {
	roomnumber int
}

// Customers defines the struct for the customers
type Customers struct {
	Customer []Customer
}

// Customer defines the struct for the  customer
type Customer struct {
	name string
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

	client, err := ent.Open("sqlite3", "file:contagion.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewRoomController(v1, client)
	controllers.NewCustomerController(v1, client)
	controllers.NewProblemController(v1, client)
	controllers.NewProblemtypeController(v1, client)

	// Set Room Data
	rooms := Rooms{
		Room: []Room{
			Room{101},
			Room{102},
		},
	}

	for _, r := range rooms.Room {
		client.Room.
			Create().
			SetROOMNUMBER(r.roomnumber).
			Save(context.Background())
	}

	// Set Customer Data
	customers := Customers{
		Customer: []Customer{
			Customer{"Kittisak Phetnae"},
			Customer{"Name Mena"},
		},
	}

	for _, n := range customers.Customer {
		client.Customer.
			Create().
			SetNAME(n.name).
			Save(context.Background())
	}
	// Set Problemtype Data
	problemtypes := Problemtypes{
		Problemtype: []Problemtype{
			Problemtype{"ห้องนั่งเล่น"},
			Problemtype{"ห้องนอน"},
			Problemtype{"ห้องน้ำ"},
		},
	}

	for _, pt := range problemtypes.Problemtype {
		client.Problemtype.
			Create().
			SetPROBLEMTYPE(pt.problemtype).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
