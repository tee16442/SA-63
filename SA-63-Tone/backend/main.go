package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/witsarut7/app/controllers"
	_ "github.com/witsarut7/app/docs"
	"github.com/witsarut7/app/ent"
)

// Customers defines the struct for the customers
type Customers struct {
	Customer []Customer
}

// Customer defines the struct for the customer
type Customer struct {
	USERNAME string
}

// Roomtypes defines the struct for the roomtypes
type Roomtypes struct {
	Roomtype []Roomtype
}

// Roomtype defines the struct for the roomtype
type Roomtype struct {
	ROOMPRICE int
}

// Paymenttypes defines the struct for the paymenttypes
type Paymenttypes struct {
	Paymenttype []Paymenttype
}

// Paymenttype defines the struct for the paymenttype
type Paymenttype struct {
	TYPE string
}

// Employees defines the struct for the employees
type Employees struct {
	Employee []Employee
}

// Employee defines the struct for the employee
type Employee struct {
	EMPLOYEENAME     string
	EMPLOYEEPASSWORD string
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

	client, err := ent.Open("sqlite3", "file:sa-g17.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewCustomerController(v1, client)
	controllers.NewRoomtypeController(v1, client)
	controllers.NewPaymenttypeController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewPaymentController(v1, client)

	// Set Customers Data
	customers := Customers{
		Customer: []Customer{
			Customer{"Tone7"},
			Customer{"Tee18"},
			Customer{"Tong"},
			Customer{"Job"},
		},
	}

	for _, cus := range customers.Customer {
		client.Customer.
			Create().
			SetUSERNAME(cus.USERNAME).
			Save(context.Background())
	}

	// Set Customers Data
	roomtypes := Roomtypes{
		Roomtype: []Roomtype{
			Roomtype{500},
			Roomtype{1000},
			Roomtype{1500},
		},
	}

	for _, rt := range roomtypes.Roomtype {
		client.Roomtype.
			Create().
			SetROOMPRICE(rt.ROOMPRICE).
			Save(context.Background())
	}

	// Set Paymenttypes Data
	paymenttypes := Paymenttypes{
		Paymenttype: []Paymenttype{
			Paymenttype{"ชำระเงินสด"},
			Paymenttype{"ชำระเงินผ่านออนไลน์"},
		},
	}

	for _, pt := range paymenttypes.Paymenttype {
		client.Paymenttype.
			Create().
			SetTYPE(pt.TYPE).
			Save(context.Background())
	}

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"Witsarut", "12345679a"},
			Employee{"Varisa", "123456789b"},
			Employee{"Varisa", "123456789c"},
		},
	}

	for _, em := range employees.Employee {
		client.Employee.
			Create().
			SetEMPLOYEENAME(em.EMPLOYEENAME).
			SetEMPLOYEEPASSWORD(em.EMPLOYEEPASSWORD).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
