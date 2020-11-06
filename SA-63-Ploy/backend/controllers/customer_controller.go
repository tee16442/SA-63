package controllers

import (
	"context"
	"strconv"

	"github.com/Sujitnapa21/app/ent"
	"github.com/Sujitnapa21/app/ent/customer"
	"github.com/gin-gonic/gin"
)




// CustomerController defines the struct for the customer controller
type CustomerController struct {
	client *ent.Client
	router gin.IRouter
}

// Customer defines the struct for the customer controller
type Customer struct {
	Email    string
	Username string
	
}

// CreateCustomer handles POST requests for adding customer entities
// @Summary Create customer
// @Description Create customer
// @ID create-customer
// @Accept   json
// @Produce  json
// @Param customer body ent.Customers true "Customer entity"
// @Success 200 {object} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers [post]
func (ctl *CustomerController) CreateCustomer(c *gin.Context) {
	obj := ent.Customer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "customer binding failed",
		})
		return
	}

	cus, err := ctl.client.Customer.
		Create().
		SetEmail(obj.Email).
		SetUsername(obj.Username).

		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cus)
}

// GetCustomer handles GET requests to retrieve a customer entity
// @Summary Get a customer entity by ID
// @Description get customers by ID
// @ID get-Customer
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers/{id} [get]
func (ctl *CustomerController) GetCustomer(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cus, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cus)
}

// ListCustomer handles request to get a list of customer entities
// @Summary List customer entities
// @Description list customer entities
// @ID list-customer
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers [get]
func (ctl *CustomerController) ListCustomer(c *gin.Context) {
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

	customers, err := ctl.client.Customer.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, customers)
}

// NewCustomerController creates and registers handles for the customer controller
func NewCustomerController(router gin.IRouter, client *ent.Client) *CustomerController {
	cusc := &CustomerController{
		client: client,
		router: router,
	}
	cusc.register()
	return cusc
}

func (ctl *CustomerController) register() {
	customers := ctl.router.Group("/customers")

	customers.GET("", ctl.ListCustomer)
	customers.POST("", ctl.CreateCustomer)
	customers.GET(":id", ctl.GetCustomer)
}