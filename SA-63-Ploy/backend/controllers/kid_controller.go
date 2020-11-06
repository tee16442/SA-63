package controllers

import (
	"context"
	"strconv"

	"github.com/Sujitnapa21/app/ent"
	"github.com/Sujitnapa21/app/ent/kid"
	"github.com/gin-gonic/gin"
)




// KidController defines the struct for the kid controller
type KidController struct {
	client *ent.Client
	router gin.IRouter
}

// Kid defines the struct for the kid controller
type Kid struct {
	Amount    int
	
}

// CreateKid handles POST requests for adding kid entities
// @Summary Create kid
// @Description Create kid
// @ID create-kid
// @Accept   json
// @Produce  json
// @Param kid body ent.Kids true "Kid entity"
// @Success 200 {object} ent.Kid
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /kids [post]
func (ctl *KidController) CreateKid(c *gin.Context) {
	obj := ent.Kid{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "kid binding failed",
		})
		return
	}

	k, err := ctl.client.Kid.
		Create().
		SetAmount(obj.Amount).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, k)
}

// GetKid handles GET requests to retrieve a kid entity
// @Summary Get a kid entity by ID
// @Description get kids by ID
// @ID get-Kid
// @Produce  json
// @Param id path int true "Kid ID"
// @Success 200 {object} ent.Kid
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /kids/{id} [get]
func (ctl *KidController) GetKid(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	k, err := ctl.client.Kid.
		Query().
		Where(kid.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, k)
}

// ListKid handles request to get a list of kid entities
// @Summary List kid entities
// @Description list kid entities
// @ID list-kid
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Kid
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /kids [get]
func (ctl *KidController) ListKid(c *gin.Context) {
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

	kids, err := ctl.client.Kid.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, kids)
}

// NewKidController creates and registers handles for the kid controller
func NewKidController(router gin.IRouter, client *ent.Client) *KidController {
	kc := &KidController{
		client: client,
		router: router,
	}
	kc.register()
	return kc
}

func (ctl *KidController) register() {
	kids := ctl.router.Group("/kids")

	kids.GET("", ctl.ListKid)
	kids.POST("", ctl.CreateKid)
	kids.GET(":id", ctl.GetKid)
}