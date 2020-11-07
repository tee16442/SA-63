package controllers
 
import (
   "context"
   "strconv"
   "github.com/chanudomsr/app/ent"
   "github.com/chanudomsr/app/ent/room"
   "github.com/gin-gonic/gin"
)
 
// RoomController defines the struct for the room controller
type RoomController struct {
   client *ent.Client
   router gin.IRouter
}
// Room defines the struct for the room 
type Room struct {
	ROOMNUMBER int
 }
// CreateRoom handles POST requests for adding room entities
// @Summary Create room
// @Description Create room
// @ID create-room
// @Accept   json
// @Produce  json
// @Param room body ent.Room true "Room entity"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Rooms [post]
func (ctl *RoomController) CreateRoom(c *gin.Context) {
	obj := ent.Room{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "room binding failed",
		})
		return
	}
  
	r, err := ctl.client.Room.
		Create().
		SetROOMNUMBER(obj.ROOMNUMBER).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, r)
 }
// GetRoom handles GET requests to retrieve a room entity
// @Summary Get a room entity by ID
// @Description get room by ID
// @ID get-room
// @Produce  json
// @Param id path int true "Room ID"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms/{id} [get]
func (ctl *RoomController) GetRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, r)
 }
 // ListRoom handles request to get a list of room entities
// @Summary List room entities
// @Description list room entities
// @ID list-room
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms [get]
func (ctl *RoomController) ListRoom(c *gin.Context) {
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
  
	rooms, err := ctl.client.Room.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, rooms)
 }
 // NewRoomController creates and registers handles for the Room controller
func NewRoomController(router gin.IRouter, client *ent.Client) *RoomController {
	rc := &RoomController{
		client: client,
		router: router,
	}
	rc.register()
	return rc
 }
  
 // InitRoomController registers routes to the main engine
 func (ctl *RoomController) register() {
	rooms := ctl.router.Group("/rooms")
  
	rooms.GET("", ctl.ListRoom)
  
	// CRUD
	rooms.POST(":id", ctl.CreateRoom)
	rooms.GET(":id", ctl.GetRoom)

	
 }
 
