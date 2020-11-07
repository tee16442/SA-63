package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/chanudomsr/app/ent"
	"github.com/chanudomsr/app/ent/books"
)

// BooksController defines the struct for the books controller
type BooksController struct {
	client *ent.Client
	router gin.IRouter
}

// Books defines the struct for the books
type Books struct {
	CHECKINDATE  string
	
}

// CreateBooks handles POST requests for adding books entities
// @Summary Create books
// @Description Create books
// @ID create-books
// @Accept   json
// @Produce  json
// @Param books body ent.Books true "Books entity"
// @Success 200 {object} ent.Books
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookss [post]
func (ctl *BooksController) CreateBooks(c *gin.Context) {
	obj := ent.Books{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "books binding failed",
		})
		return
	}

	b, err := ctl.client.Books.
		Create().
		SetCHECKINDATE(obj.CHECKINDATE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// GetBooks handles GET requests to retrieve a books entity
// @Summary Get a books entity by ID
// @Description get books by ID
// @ID get-books
// @Produce  json
// @Param id path int true "Books ID"
// @Success 200 {object} ent.Books
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookss/{id} [get]
func (ctl *BooksController) GetBooks(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := ctl.client.Books.
		Query().
		Where(books.IDEQ(int(id))).
		Only(context.Background())
		
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
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
// @Router /books [get]
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

	Bookss, err := ctl.client.Books.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Bookss)
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

// InitBooksController registers routes to the main engine
func (ctl *BooksController) register() {
	bookss := ctl.router.Group("/bookss")

	bookss.GET("", ctl.ListBooks)

	// CRUD
	bookss.POST("", ctl.CreateBooks)
	bookss.GET(":id", ctl.GetBooks)

}
