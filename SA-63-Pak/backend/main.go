package main
 
import (
   "context"
   "log"
 
   "github.com/PBank/app/controllers"
   _ "github.com/PBank/app/docs"
   "github.com/PBank/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)
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
type Roomtypes struct {
	Roomtype []Roomtype
}
type Roomtype struct {
	ROOMPRICE int
	TYPEDETAIL string
}
type Roomstates struct {
	Roomstate []Roomstate
}
type Roomstate struct {
	ROOMSTATE  string
}
type Promotions struct {
	Promotion []Promotion
}
type Promotion struct {
	PROMOTIONDETAIL  string
}
func main() {
	router := gin.Default()
	router.Use(cors.Default())
  
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()
  
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
  
	v1 := router.Group("/api/v1")
	controllers.NewRoomController(v1, client)
	controllers.NewPromotionController(v1, client)
	controllers.NewRoomstateController(v1, client)
	controllers.NewRoomtypeController(v1, client)
  
	

	
	// Set RoomTypes Data
	roomtypes := Roomtypes{
		Roomtype: []Roomtype{
			Roomtype{500, "Standard"},
			Roomtype{1000, "Superior"},
			Roomtype{1500, "Deluxe"},
		},
	}

	for _, rt := range roomtypes.Roomtype {
		client.Roomtype.
			Create().
			
			SetROOMPRICE(rt.ROOMPRICE).
			SetTYPEDEATAIL(rt.TYPEDETAIL).
			Save(context.Background())
	}
	// Set Promotions Data
	promotions := Promotions{
		Promotion: []Promotion{
			Promotion{ "ไม่มีโปรโมชั่น"},
			Promotion{ "ลดราคา50%"},
			Promotion{ "จอง1ห้องพักฟรีเพิ่มอีก1ห้อง"},
		},
	}

	for _, pm := range promotions.Promotion {
		client.Promotion.
			Create().
			
			SetPROMOTIONDETAIL(pm.PROMOTIONDETAIL).
			Save(context.Background())
	}
	// Set RoomStates Data
	roomstates := Roomstates{
		Roomstate: []Roomstate{
			Roomstate{"Availability"},
			Roomstate{"No Availability"},
		},
	}

	for _, rs := range roomstates.Roomstate {
		client.Roomstate.
			Create().
			
			SetROOMSTATE(rs.ROOMSTATE).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
 }
 

 