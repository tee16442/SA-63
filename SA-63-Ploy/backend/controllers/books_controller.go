package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/Sujitnapa21/app/ent"
	"github.com/Sujitnapa21/app/ent/adult"
	"github.com/Sujitnapa21/app/ent/customer"
	"github.com/Sujitnapa21/app/ent/kid"
	"github.com/Sujitnapa21/app/ent/room"
	"github.com/Sujitnapa21/app/ent/roomamount"

	"github.com/gin-gonic/gin"
)

// BooksController defines the struct for the books controller
type BooksController struct {
	client *ent.Client
	router gin.IRouter
}

// Books  defines the struct for the books controller
type Books struct {
	Customer   int
	Room       int
	Adult      int
	Kid        int
	Roomamount int
	Checkin    string
	Checkout   string
}

// CreateBooks handles POST requests for adding books entities
// @Summary Create books
// @Description Create books
// @ID create-books
// @Accept   json
// @Produce  json
// @Param books body ent.Books true "Books entity"
// @Success 200 {object} ent.Books
// @Failure 400 {object} gin.H.
// @Failure 500 {object} gin.H
// @Router /bookss [post]
func (ctl *BooksController) CreateBooks(c *gin.Context) {
	obj := Books{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "books binding failed",
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

	a, err := ctl.client.Adult.
		Query().
		Where(adult.IDEQ(int(obj.Adult))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "adult not found",
		})
		return
	}

	k, err := ctl.client.Kid.
		Query().
		Where(kid.IDEQ(int(obj.Kid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "kid not found",
		})
		return
	}

	ra, err := ctl.client.Roomamount.
		Query().
		Where(roomamount.IDEQ(int(obj.Roomamount))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "roomamount not found",
		})
		return
	}

	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room found",
		})
		return
	}

	checkin, err := time.Parse(time.RFC3339, obj.Checkin)
	checkout, err := time.Parse(time.RFC3339, obj.Checkout)

	b, err := ctl.client.Books.
		Create().
		SetCustomer(cus).
		SetRoom(r).
		SetAdult(a).
		SetKid(k).
		SetRoomamount(ra).
		SetCheckin(checkin).
		SetCheckout(checkout).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// ListBooks handles request to get a list of books entities
// @Summary List books entities
// @Description list books entities
// @ID list-books
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Books 
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookss [get]
func (ctl *BooksController) ListBooks(c *gin.Context) {
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

	bookss, err := ctl.client.Books.
		Query().
		WithCustomer().
		WithRoom().
		WithAdult().
		WithKid().
		WithRoomamount().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bookss)
}

// NewBooksController creates and registers handles for the books controller
func NewBooksController(router gin.IRouter, client *ent.Client) *BooksController {
	bc := &BooksController{
		client: client,
		router: router,
	}
	bc.register()
	return bc
}

func (ctl *BooksController) register() {
	bookss := ctl.router.Group("/bookss")
	bookss.POST("", ctl.CreateBooks)
	bookss.GET("", ctl.ListBooks)

}
