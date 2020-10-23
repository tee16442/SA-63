package main

import (
	"context"
	"log"

	"github.com/ADMIN/app/controllers"
	_ "github.com/ADMIN/app/docs"
	"github.com/ADMIN/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Titles  defines the struct for the titles
//---------------------------------------------
type Titles struct {
	Title []Title
}

// Title  defines the struct for the title
type Title struct {
	titletype string
}

// Genders  defines the struct for the genders
//---------------------------------------------
type Genders struct {
	Gender []Gender
}

// Gender  defines the struct for the gender
type Gender struct {
	gender string
}

// Customertypes  defines the struct for the Customertypes
//-------------------------------------------------------
type Customertypes struct {
	Customertype []Customertype
}

// Customertype  defines the struct for the Customertype
type Customertype struct {
	customertype string
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
	controllers.NewCustomerController(v1, client)
	controllers.NewCustomertypeController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewTitleController(v1, client)

	// Set Titles Data
	titles := Titles{
		Title: []Title{
			Title{"นาย"},
			Title{"นาง"},
			Title{"นางสาว"},
		},
	}

	for _, ti := range titles.Title {
		client.Title.
			Create().
			SetTITLETYPE(ti.titletype).
			Save(context.Background())
	}

	// Set Genders Data
	genders := Genders{
		Gender: []Gender{
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}

	for _, g := range genders.Gender {
		client.Gender.
			Create().
			SetGENDER(g.gender).
			Save(context.Background())
	}

	// Set Customertypes Data
	customertypes := Customertypes{
		Customertype: []Customertype{
			Customertype{"ทั่วไป"},
			Customertype{"VIP"},
		},
	}

	for _, ct := range customertypes.Customertype {
		client.Customertype.
			Create().
			SetCUSTOMERTYPE(ct.customertype).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
