package controllers

import (
	"context"
	"strconv"

	"github.com/Sujitnapa21/app/ent"
	"github.com/Sujitnapa21/app/ent/adult"
	"github.com/gin-gonic/gin"
)

// AdultController defines the struct for the adult controller
type AdultController struct {
	client *ent.Client
	router gin.IRouter
}

// Adult defines the struct for the adult controller
type Adult struct {
	Amount int
}

// CreateAdult handles POST requests for adding adult entities
// @Summary Create adult
// @Description Create adult
// @ID create-adult
// @Accept   json
// @Produce  json
// @Param adult body ent.Adults true "Adult entity"
// @Success 200 {object} ent.Adult
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /adults [post]
func (ctl *AdultController) CreateAdult(c *gin.Context) {
	obj := ent.Adult{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "adult binding failed",
		})
		return
	}

	a, err := ctl.client.Adult.
		Create().
		SetAmount(obj.Amount).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, a)
}

// GetAdult handles GET requests to retrieve a adult entity
// @Summary Get a adult entity by ID
// @Description get adults by ID
// @ID get-Adult
// @Produce  json
// @Param id path int true "Adult ID"
// @Success 200 {object} ent.Adult
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /adults/{id} [get]
func (ctl *AdultController) GetAdult(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	a, err := ctl.client.Adult.
		Query().
		Where(adult.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, a)
}

// ListAdult handles request to get a list of adult entities
// @Summary List adult entities
// @Description list adult entities
// @ID list-adult
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Adult
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /adults [get]
func (ctl *AdultController) ListAdult(c *gin.Context) {
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

	adults, err := ctl.client.Adult.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, adults)
}

// NewAdultController creates and registers handles for the adult controller
func NewAdultController(router gin.IRouter, client *ent.Client) *AdultController {
	ac := &AdultController{
		client: client,
		router: router,
	}
	ac.register()
	return ac
}

func (ctl *AdultController) register() {
	adults := ctl.router.Group("/adults")

	adults.GET("", ctl.ListAdult)
	adults.POST("", ctl.CreateAdult)
	adults.GET(":id", ctl.GetAdult)
}
