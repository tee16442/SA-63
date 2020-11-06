package controllers

import (
	"context"
	"strconv"

	"github.com/Sujitnapa21/app/ent"
	"github.com/Sujitnapa21/app/ent/roomamount"
	"github.com/gin-gonic/gin"
)




// RoomamountController defines the struct for the roomamount controller
type RoomamountController struct {
	client *ent.Client
	router gin.IRouter
}

// Roomamount defines the struct for the roomamount controller
type Roomamount struct {
	Amount string
	
}

// CreateRoomamount handles POST requests for adding roomamount entities
// @Summary Create roomamount
// @Description Create roomamount
// @ID create-roomamount
// @Accept   json
// @Produce  json
// @Param roomamount body ent.Roomamount true "Roomamount entity"
// @Success 200 {object} ent.Roomamount
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomamounts [post]
func (ctl *RoomamountController) CreateRoomamount(c *gin.Context) {
	obj := ent.Roomamount{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "roomamount binding failed",
		})
		return
	}

	ra, err := ctl.client.Roomamount.
		Create().
		SetAmount(obj.Amount).

		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ra)
}

// GetRoomamount handles GET requests to retrieve a roomamount entity
// @Summary Get a roomamount entity by ID
// @Description get roomamounts by ID
// @ID get-RoomAmount
// @Produce  json
// @Param id path int true "Roomamount ID"
// @Success 200 {object} ent.Roomamount
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomamounts/{id} [get]
func (ctl *RoomamountController) GetRoomamount(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ra, err := ctl.client.Roomamount.
		Query().
		Where(roomamount.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ra)
}

// ListRoomamount handles request to get a list of roomamount entities
// @Summary List roomamount entities
// @Description list roomamount entities
// @ID list-roomamount
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Roomamount
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomamounts [get]
func (ctl *RoomamountController) ListRoomamount(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	roomamounts, err := ctl.client.Roomamount.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, roomamounts)
}

// NewRoomamountController creates and registers handles for the roomamount controller
func NewRoomamountController(router gin.IRouter, client *ent.Client) *RoomamountController {
	rac := &RoomamountController{
		client: client,
		router: router,
	}
	rac.register()
	return rac
}

func (ctl *RoomamountController) register() {
	roomamounts := ctl.router.Group("/roomamounts")

	roomamounts.GET("", ctl.ListRoomamount)
	roomamounts.POST("", ctl.CreateRoomamount)
	roomamounts.GET(":id", ctl.GetRoomamount)
}