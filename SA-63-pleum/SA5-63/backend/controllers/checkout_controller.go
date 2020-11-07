package controllers
 
import (
   "context"
   "strconv"
   "time"
   "github.com/chanudomsr/app/ent"
   "github.com/chanudomsr/app/ent/checkout"
   "github.com/gin-gonic/gin"
   "github.com/chanudomsr/app/ent/books"
   "github.com/chanudomsr/app/ent/customer"
   "github.com/chanudomsr/app/ent/employee"
   "github.com/chanudomsr/app/ent/room"
)
 
// CheckoutController defines the struct for the checkout controller
type CheckoutController struct {
   client *ent.Client
   router gin.IRouter
}

// Checkout defines the struct for the checkout 
type Checkout struct {
	Books        int
	Employee     int
	Customer     int
	Room         int
	Checkoutdate string
 }
 
// CreateCheckout handles POST requests for adding checkout entities
// @Summary Create checkout
// @Description Create checkout
// @ID create-checkout
// @Accept   json
// @Produce  json
// @Param checkout body Checkout true "Checkout entity"
// @Success 200 {object} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts [post]
func (ctl *CheckoutController) CreateCheckout(c *gin.Context) {
	obj := Checkout{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "checkout binding failed",
		})
		return
	}
	
	b, err := ctl.client.Books.
		Query().
		Where(books.IDEQ(int(obj.Books))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "books not found",
		})
		return
	}

	cus, err := ctl.client.Customer.
		Query().
		Where(customer.IDEQ(int(obj.Customer))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "customer not found",
		})
		return
	}

	em, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.Employee))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}

	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room not found",
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.Checkoutdate)
	co, err := ctl.client.Checkout.
		Create().
		SetCHECKOUTDATE(time).
		SetBooks(b).
		SetCustomer(cus).
		SetEmployee(em).
		SetRoom(r).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, co)
 }
 // GetCheckout handles GET requests to retrieve a checkout entity
// @Summary Get a checkout entity by ID
// @Description get checkout by ID
// @ID get-checkout
// @Produce  json
// @Param id path int true "Checkout ID"
// @Success 200 {object} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts/{id} [get]
func (ctl *CheckoutController) GetCheckout(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	co, err := ctl.client.Checkout.
		Query().
		WithCustomer().
		WithEmployee().
		WithRoom().
		Where(checkout.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, co)
 }

// ListCheckout handles request to get a list of checkout entities
// @Summary List checkout entities
// @Description list checkout entities
// @ID list-checkout
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Checkout
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /checkouts [get]
func (ctl *CheckoutController) ListCheckout(c *gin.Context) {
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
  
	checkouts, err := ctl.client.Checkout.
		Query().
		//WithBooks().
		WithCustomer().
		WithEmployee().
		WithRoom().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, checkouts)
 }

// NewCheckoutController creates and registers handles for the checkout controller
func NewCheckoutController(router gin.IRouter, client *ent.Client) *CheckoutController {
	coc := &CheckoutController{
		client: client,
		router: router,
	}
	coc.register()
	return coc
 }
  
 // InitCheckoutController registers routes to the main engine
 func (ctl *CheckoutController) register() {
	checkouts := ctl.router.Group("/checkouts")
  
	checkouts.GET("", ctl.ListCheckout)
  
	// CRUD
	checkouts.POST("", ctl.CreateCheckout)
	checkouts.GET(":id", ctl.GetCheckout)
	
 }
	