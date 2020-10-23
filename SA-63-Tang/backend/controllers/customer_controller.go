package controllers

import (
	"context"

	"github.com/ADMIN/app/ent"
	"github.com/ADMIN/app/ent/customertype"
	"github.com/ADMIN/app/ent/gender"
	"github.com/ADMIN/app/ent/title"
	"github.com/gin-gonic/gin"
)

// CustomerController defines the struct for the customer controller
type CustomerController struct {
	client *ent.Client
	router gin.IRouter
}

// Customer defines the struct for the customer
type Customer struct {
	Name         string
	Age          int
	Email        string
	Password     string
	Username     string
	Gender       int
	Customertype int
	Title        int
}

// CreateCustomer handles POST requests for adding customer entities
// @Summary Create customer
// @Description Create customer
// @ID create-customer
// @Accept   json
// @Produce  json
// @Param customer body Customer true "customer entity"
// @Success 200 {object} ent.Customer
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /customers [post]
func (ctl *CustomerController) CreateCustomer(c *gin.Context) {
	obj := Customer{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "customer binding failed",
		})
		return
	}
	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "owner not found",
		})
		return
	}
	ct, err := ctl.client.Customertype.
		Query().
		Where(customertype.IDEQ(int(obj.Customertype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "owner not found",
		})
		return
	}
	ti, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(int(obj.Title))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "owner not found",
		})
		return
	}

	cus, err := ctl.client.Customer.
		Create().
		SetNAME(obj.Name).
		SetAGE(obj.Age).
		SetEMAIL(obj.Email).
		SetPASSWORD(obj.Password).
		SetUSERNAME(obj.Username).
		SetGender(g).
		SetCustomertype(ct).
		SetTitle(ti).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cus)
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

// InitCustomerController registers routes to the main engine
func (ctl *CustomerController) register() {
	customers := ctl.router.Group("/customers")



	// CRUD
	customers.POST("", ctl.CreateCustomer)
	

}
