package controllers
 
import (
   "context"
   "strconv"
   "github.com/PBank/app/ent"
   "github.com/PBank/app/ent/promotion"
   "github.com/gin-gonic/gin"
)
 
// PromotionController defines the struct for the promotion controller
type PromotionController struct {
   client *ent.Client
   router gin.IRouter
}

type Promotion struct {
	PROMOTION   string
	PROMOTIONDETAIL  string
}

// CreatePromotion handles POST requests for adding promotion entities
// @Summary Create promotion
// @Description Create promotion
// @ID create-promotion
// @Accept   json
// @Produce  json
// @Param promotion body ent.Promotion true "Promotion entity"
// @Success 200 {object} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions [post]
func (ctl *PromotionController) CreatePromotion(c *gin.Context) {
	obj := ent.Promotion{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "promotion binding failed",
		})
		return
	}
  
	pm, err := ctl.client.Promotion.
		Create().
		SetPROMOTIONDETAIL(obj.PROMOTIONDETAIL).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, pm)
 }
 // GetPromotion handles GET requests to retrieve a promotion entity
// @Summary Get a promotion entity by ID
// @Description get promotion by ID
// @ID get-promotion
// @Produce  json
// @Param id path int true "Promotion ID"
// @Success 200 {object} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions/{id} [get]
func (ctl *PromotionController) GetPromotion(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	pm, err := ctl.client.Promotion.
		Query().
		Where(promotion.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, pm)
 }
 // ListPromotion handles request to get a list of promotion entities
// @Summary List promotion entities
// @Description list promotion entities
// @ID list-promotion
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Promotion
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /promotions [get]
func (ctl *PromotionController) ListPromotion(c *gin.Context) {
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
  
	promotions, err := ctl.client.Promotion.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, promotions)
 }

 // NewPromotionController creates and registers handles for the promotion controller
func NewPromotionController(router gin.IRouter, client *ent.Client) *PromotionController {
	pmc := &PromotionController{
		client: client,
		router: router,
	}
	pmc.register()
	return pmc
 }
  
 // InitPromotionController registers routes to the main engine
 func (ctl *PromotionController) register() {
	promotions := ctl.router.Group("/promotions")
  
	promotions.GET("", ctl.ListPromotion)
	// CRUD
	promotions.POST("", ctl.CreatePromotion)
	promotions.GET(":id", ctl.GetPromotion)
 }
 