package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tee16/app/ent"
	"github.com/tee16/app/ent/customer"
	"github.com/tee16/app/ent/problem"
	"github.com/tee16/app/ent/problemtype"
	"github.com/tee16/app/ent/room"
)

//ProblemController defines the struct for the problem controller
type ProblemController struct {
	client *ent.Client
	router gin.IRouter
}

//Problem defines the struct for the problem
type Problem struct {
	Problemdetail string
	Problemtype   int
	Name          int
	Roomnumber    int
}

// CreateProblem handles POST requests for adding problem entities
// @Summary Create problem
// @Description Create problem
// @ID create-problem
// @Accept   json
// @Produce  json
// @Param problem body Problem true "problem entity"
// @Success 200 {object} ent.Problem
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problems [post]
func (ctl *ProblemController) CreateProblem(c *gin.Context) {
	obj := Problem{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Problem binding failed",
		})
		return
	}
	pt, err := ctl.client.Problemtype.
		Query().
		Where(problemtype.IDEQ(int(obj.Problemtype))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "owner not found",
		})
		return
	}
	n, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.Name))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "owner not found",
		})
		return
	}
	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Roomnumber))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "owner not found",
		})
		return
	}

	u, err := ctl.client.Problem.
		Create().
		SetPROBLEMDETAIL(obj.Problemdetail).
		SetProblemtype(pt).
		SetCustomer(n).
		SetRoom(r).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetProblem handles GET requests to retrieve a problem entity
// @Summary Get a problem entity by ID
// @Description get problem by ID
// @ID get-problem
// @Produce  json
// @Param id path int true "problem ID"
// @Success 200 {object} ent.Problem
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problems/{id} [get]
func (ctl *ProblemController) GetProblem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Problem.
		Query().
		Where(problem.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListProblem handles request to get a list of problem entities
// @Summary List problem entities
// @Description list problem entities
// @ID list-problem
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Problem
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problems [get]
func (ctl *ProblemController) ListProblem(c *gin.Context) {
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

	problems, err := ctl.client.Problem.
		Query().WithCustomer().WithProblemtype().WithRoom().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, problems)
}

// NewProblemController creates and registers handles for the problem controller
func NewProblemController(router gin.IRouter, client *ent.Client) *ProblemController {
	uc := &ProblemController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitProblemController registers routes to the main engine
func (ctl *ProblemController) register() {
	problems := ctl.router.Group("/problems")

	problems.GET("", ctl.ListProblem)

	// CRUD
	problems.POST("", ctl.CreateProblem)
	problems.GET(":id", ctl.GetProblem)

}
