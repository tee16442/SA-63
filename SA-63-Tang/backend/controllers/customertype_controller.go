package controllers

import (
	"context"
	"strconv"

	"github.com/ADMIN/app/ent"
	"github.com/ADMIN/app/ent/customertype"
	"github.com/gin-gonic/gin"
)

// CustomertypeController defines the struct for the customertype controller
type CustomertypeController struct {
	client *ent.Client
	router gin.IRouter
}

// Customertype defines the struct for the Customertype
type Customertype struct {
	Customertype    string
}
// CreateCustomertype handles POST requests for adding customertype entities
// @Summary Create customertype
// @Description Create customertype
// @ID create-customertype
// @Accept   json
// @Produce  json
// @Param customertype body ent.Customertype true "Customertype entity"
// @Success 200 {object} ent.Customertype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customertypes [post]
func (ctl *CustomertypeController) CreateCustomertype(c *gin.Context) {
	obj := ent.Customertype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "customertype binding failed",
		})
		return
	}

	u, err := ctl.client.Customertype.
		Create().
		SetCUSTOMERTYPE(obj.CUSTOMERTYPE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetCustomertype handles GET requests to retrieve a customertype entity
// @Summary Get a customertype entity by ID
// @Description get customertype by ID
// @ID get-customertype
// @Produce  json
// @Param id path int true "Customertype ID"
// @Success 200 {object} ent.Customertype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customertypes/{id} [get]
func (ctl *CustomertypeController) GetCustomertype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Customertype.
		Query().
		Where(customertype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListCustomertype handles request to get a list of customertype entities
// @Summary List customertype entities
// @Description list customertype entities
// @ID list-customertype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Customertype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customertypes [get]
func (ctl *CustomertypeController) ListCustomertype(c *gin.Context) {
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

	customertypes, err := ctl.client.Customertype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, customertypes)
}

// NewCustomertypeController creates and registers handles for the customertype controller
func NewCustomertypeController(router gin.IRouter, client *ent.Client) *CustomertypeController {
	uc := &CustomertypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitCustomertypeController registers routes to the main engine
func (ctl *CustomertypeController) register() {
	customertypes := ctl.router.Group("/customertypes")
	customertypes.GET("", ctl.ListCustomertype)

	// CRUD
	customertypes.POST("", ctl.CreateCustomertype)
	customertypes.GET(":id", ctl.GetCustomertype)
}
