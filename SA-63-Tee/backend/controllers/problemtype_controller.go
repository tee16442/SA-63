package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tee16/app/ent"
	"github.com/tee16/app/ent/problemtype"
)

// ProblemtypeController defines the struct for the problemtype controller
type ProblemtypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateProblemtype handles POST requests for adding problemtype entities
// @Summary Create problemtype
// @Description Create problemtype
// @ID create-problemtype
// @Accept   json
// @Produce  json
// @Param problemtype body ent.Problemtype true "problemtype entity"
// @Success 200 {object} ent.Problemtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemtypes [post]
func (ctl *ProblemtypeController) CreateProblemtype(c *gin.Context) {
	obj := ent.Problemtype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "ProblemType binding failed",
		})
		return
	}

	u, err := ctl.client.Problemtype.
		Create().
		SetPROBLEMTYPE(obj.PROBLEMTYPE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetProblemtype handles GET requests to retrieve a problemtype entity
// @Summary Get a problemtype entity by ID
// @Description get problemtype by ID
// @ID get-problemtype
// @Produce  json
// @Param id path int true "problemtype ID"
// @Success 200 {object} ent.Problemtype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemtypes/{id} [get]
func (ctl *ProblemtypeController) GetProblemtype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Problemtype.
		Query().
		Where(problemtype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListProblemtype handles request to get a list of problemtype entities
// @Summary List problemtype entities
// @Description list problemtype entities
// @ID list-problemtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Problemtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemtypes [get]
func (ctl *ProblemtypeController) ListProblemtype(c *gin.Context) {
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

	problemtypes, err := ctl.client.Problemtype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, problemtypes)
}

// NewProblemtypeController creates and registers handles for the problemtype controller
func NewProblemtypeController(router gin.IRouter, client *ent.Client) *ProblemtypeController {
	uc := &ProblemtypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitProblemtypeController registers routes to the main engine
func (ctl *ProblemtypeController) register() {
	problemtypes := ctl.router.Group("/problemtypes")

	problemtypes.GET("", ctl.ListProblemtype)

	// CRUD
	problemtypes.POST("", ctl.CreateProblemtype)
	problemtypes.GET(":id", ctl.GetProblemtype)

}
