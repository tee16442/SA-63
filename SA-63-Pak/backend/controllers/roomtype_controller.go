package controllers
 
import (
   "context"
   "strconv"
   "github.com/PBank/app/ent"
   "github.com/PBank/app/ent/roomtype"
   "github.com/gin-gonic/gin"
)
 
// RoomtypeController defines the struct for the roomtype controller
type RoomtypeController struct {
   client *ent.Client
   router gin.IRouter
}

type Roomtype struct {
	ROOMTYPE  string
	ROOMPRICE string
	TYPEDETAIL string
}

// CreateRoomtype handles POST requests for adding roomtype entities
// @Summary Create roomtype
// @Description Create roomtype
// @ID create-roomtype
// @Accept   json
// @Produce  json
// @Param roomtype body ent.Roomtype true "Roomtype entity"
// @Success 200 {object} ent.Roomtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomtypes [post]
func (ctl *RoomtypeController) CreateRoomtype(c *gin.Context) {
	obj := ent.Roomtype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "roomtype binding failed",
		})
		return
	}
  
	rt, err := ctl.client.Roomtype.
		Create().
		SetROOMPRICE(obj.ROOMPRICE).
		SetTYPEDEATAIL(obj.TYPEDEATAIL).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, rt)
 }
 // GetRoomtype handles GET requests to retrieve a roomtype entity
// @Summary Get a roomtype entity by ID
// @Description get roomtype by ID
// @ID get-roomtype
// @Produce  json
// @Param id path int true "Roomtype ID"
// @Success 200 {object} ent.Roomtype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomtypes/{id} [get]
func (ctl *RoomtypeController) GetRoomtype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	rt, err := ctl.client.Roomtype.
		Query().
		Where(roomtype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, rt)
 }
 // ListRoomtype handles request to get a list of roomtype entities
// @Summary List roomtype entities
// @Description list roomtype entities
// @ID list-roomtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Roomtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomtypes [get]
func (ctl *RoomtypeController) ListRoomtype(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	roomtypes, err := ctl.client.Roomtype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, roomtypes)
 }

 // NewRoomtypeController creates and registers handles for the roomtype controller
func NewRoomtypeController(router gin.IRouter, client *ent.Client) *RoomtypeController {
	rtc := &RoomtypeController{
		client: client,
		router: router,
	}
	rtc.register()
	return rtc
 }
 // InitRoomtypeController registers routes to the main engine
 func (ctl *RoomtypeController) register() {
	roomtypes := ctl.router.Group("/roomtypes")
  
	roomtypes.GET("", ctl.ListRoomtype)
  
	// CRUD
	roomtypes.POST("", ctl.CreateRoomtype)
	roomtypes.GET(":id", ctl.GetRoomtype)
 }
 