package controllers
 
import (
   "context"
   "strconv"
   "github.com/PBank/app/ent"
   "github.com/PBank/app/ent/roomstate"
   "github.com/gin-gonic/gin"
)
 
// RoomstateController defines the struct for the roomstate controller
type RoomstateController struct {
   client *ent.Client
   router gin.IRouter
}

type Roomstate struct {
	STATE   string
	ROOMSTATE  string
}

// CreateRoomstate handles POST requests for adding roomstate entities
// @Summary Create roomstate
// @Description Create roomstate
// @ID create-roomstate
// @Accept   json
// @Produce  json
// @Param roomstate body ent.Roomstate true "Roomstate entity"
// @Success 200 {object} ent.Roomstate
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstates [post]
func (ctl *RoomstateController) CreateRoomstate(c *gin.Context) {
	obj := ent.Roomstate{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "roomstate binding failed",
		})
		return
	}
  
	rs, err := ctl.client.Roomstate.
		Create().
		
		SetROOMSTATE(obj.ROOMSTATE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, rs)
 }
 // GetRoomstate handles GET requests to retrieve a roomstate entity
// @Summary Get a roomstate entity by ID
// @Description get roomstate by ID
// @ID get-roomstate
// @Produce  json
// @Param id path int true "Roomstate ID"
// @Success 200 {object} ent.Roomstate
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstates/{id} [get]
func (ctl *RoomstateController) GetRoomstate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	rs, err := ctl.client.Roomstate.
		Query().
		Where(roomstate.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, rs)
 }
 // ListRoomstate handles request to get a list of roomstate entities
// @Summary List roomstate entities
// @Description list roomstate entities
// @ID list-roomstate
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Roomstate
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstates [get]
func (ctl *RoomstateController) ListRoomstate(c *gin.Context) {
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
  
	roomstates, err := ctl.client.Roomstate.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, roomstates)
 }
 // UpdateRoomstate handles PUT requests to update a roomstate entity
// @Summary Update a roomstate entity by ID
// @Description update roomstate by ID
// @ID update-roomstate
// @Accept   json
// @Produce  json
// @Param id path int true "Roomstate ID"
// @Param roomstate body ent.Roomstate true "Roomstate entity"
// @Success 200 {object} ent.Roomstate
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstates/{id} [put]
func (ctl *RoomstateController) UpdateRoomstate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Roomstate{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "roomstate binding failed",
		})
		return
	}
	obj.ID = int(id)
	rs, err := ctl.client.Roomstate.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, rs)
 }
 // NewRoomstateController creates and registers handles for the roomstate controller
func NewRoomstateController(router gin.IRouter, client *ent.Client) *RoomstateController {
	rsc := &RoomstateController{
		client: client,
		router: router,
	}
	rsc.register()
	return rsc
 }
  
 // InitRoomstateController registers routes to the main engine
 func (ctl *RoomstateController) register() {
	roomstates := ctl.router.Group("/roomstates")
  
	roomstates.GET("", ctl.ListRoomstate)
  
	// CRUD
	roomstates.POST("", ctl.CreateRoomstate)
	roomstates.GET(":id", ctl.GetRoomstate)
	roomstates.PUT(":id", ctl.UpdateRoomstate)
 }
 