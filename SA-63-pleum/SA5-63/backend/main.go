package main

import (
  "context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/chanudomsr/app/controllers"
	_ "github.com/chanudomsr/app/docs"
	"github.com/chanudomsr/app/ent"
)

// Bookss  defines the struct for the books
type Bookss struct {
	Books []Books
}

// Books  defines the struct for the books
type Books struct {
	CHECKINDATE  string
}

// Customers  defines the struct for the customers
type Customers struct {
	Customer []Customer
}

// Customer  defines the struct for the customer
type Customer struct {
	NAME string
}

// Employees  defines the struct for the employees
type Employees struct {
	Employee []Employee
}

// Employee  defines the struct for the employee
type Employee struct {
	EMPLOYEENAME     string
	EMPLOYEEPASSWORD string
}

// Rooms  defines the struct for the rooms
type Rooms struct {
	Room []Room
}

// Room  defines the struct for the room
type Room struct {
	ROOMNUMBER int
}

// @title SUT SA Example API
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

	client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewCheckoutController(v1, client)
	controllers.NewCustomerController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewRoomController(v1, client)
	// Set Customers Data
	customers := Customers{
		Customer: []Customer{
			Customer{"TONE"},
			Customer{"TANG"},
			Customer{"TEE"},
		},
	}

	for _, cus := range customers.Customer {
		client.Customer.
			Create().
			SetNAME(cus.NAME).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"EM1", "asd789"},
			Employee{"EM2", "aaa333"},
			Employee{"EM3", "eee456"},
		},
	}

	for _, em := range employees.Employee {
		client.Employee.
			Create().
			SetEMPLOYEENAME(em.EMPLOYEENAME).
			SetEMPLOYEEPASSWORD(em.EMPLOYEEPASSWORD).
			Save(context.Background())
	}

	// Set Bookss Data
	bookss := Bookss{
		Books: []Books{
			Books{"22-11-2020"},
			Books{"02-12-2020"},
			Books{"20-12-2020"},
		},
	}

	for _, b := range bookss.Books {
		client.Books.
			Create().
			SetCHECKINDATE(b.CHECKINDATE).
			Save(context.Background())
	}

	// Set Rooms Data
	rooms := Rooms{
		Room: []Room{
			Room{101},
			Room{102},
			Room{103},
			Room{104},
		},
	}

	for _, r := range rooms.Room {
		client.Room.
			Create().
			SetROOMNUMBER(r.ROOMNUMBER).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
