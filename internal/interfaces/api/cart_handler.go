package api

import (
	"strconv"

	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/dto"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartApp application.CartAppInterface
}

func NewCartHandler(cartApp application.CartAppInterface) *CartHandler {
	return &CartHandler{
		cartApp: cartApp,
	}
}

// @Summary Apply promotion
// @Description Apply a promotion to a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param cartId path int true "Cart ID"
// @Param promotionId path int true "Promotion ID"
// @Success 200 {string} string "Promotion applied successfully"
// @Failure 409 {string} string "Promotion already applied"
// @Router /api/v1/cart/{cartId}/promotion/{promotionId} [post]
func (h *CartHandler) ApplyPromotion(c *gin.Context) {
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	promotionId, err := strconv.Atoi(c.Param("promotionId"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	cart, promErr := h.cartApp.ApplyPromotion(cartId, promotionId)
	if promErr != nil {
		c.JSON(409, gin.H{"message": promErr.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Promotion applied successfully",
		"data":    cart,
	})
}

// @Summary Display cart
// @Description Display a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param cartId path int true "Cart ID"
// @Success 200 {object} string "Cart displayed successfully"
// @Router /api/v1/cart/{cartId} [get]
func (h *CartHandler) Display(c *gin.Context) {
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	cart, err := h.cartApp.Display(cartId)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Cart displayed successfully",
		"data":    cart,
	})
}

// @Summary Reset cart
// @Description Reset a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param cartId path int true "Cart ID"
// @Success 200 {string} string "Cart reset successfully"
// @Router /api/v1/cart/{cartId} [delete]
func (h *CartHandler) ResetCart(c *gin.Context) {
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	cart, resetErr := h.cartApp.ResetCart(cartId)
	if resetErr != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Cart reset successfully",
		"data":    cart,
	})
}

// @Summary Add item
// @Description Add an item to a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param cartId path int true "Cart ID"
// @Param item body dto.ItemCreateRequest true "Item"
// @Success 200 {string} string "Item added successfully"
// @Router /api/v1/cart/{cartId}/item [post]
func (h *CartHandler) AddItem(c *gin.Context) {
	var data dto.ItemCreateRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	itemEntity := data.ToEntity()
	itemEntity.CartID = cartId
	err = itemEntity.Validate()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	item, addItemErr := h.cartApp.AddItem(cartId, itemEntity)
	if addItemErr != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Item added successfully", "data": item})
}

// @Summary Add vas item
// @Description Add a vas item to a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param item body dto.VasItemCreateRequest true "VasItem"
// @Router /api/v1/cart/{cartId}/item/{itemId}/vas-item/{vasItemId} [post]
func (h *CartHandler) AddVasItem(c *gin.Context) {
	var data dto.VasItemCreateRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	vasItemEntity := data.ToEntity()
	err := vasItemEntity.Validate()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	vasItem, err := h.cartApp.AddVasItem(vasItemEntity)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Vas item added successfully", "data": vasItem})
}

// @Summary Remove Item from cart
// @Description Remove an item from a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param itemId path int true "Item ID"
// @Success 200 {string} string "Item removed successfully"
// @Router /api/v1/cart/item/{itemId} [delete]
func (h *CartHandler) RemoveItem(c *gin.Context) {

	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	cart, err := h.cartApp.RemoveItem(itemId)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Item removed successfully",
		"data":    cart,
	})
}

// @Summary Remove vas item from cart
// @Description Remove a vas item from a cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param vasItemId path int true "Vas Item ID"
// @Success 200 {string} string "Vas item removed successfully"
// @Router /api/v1/cart/vas-item/{vasItemId} [delete]
func (h *CartHandler) RemoveVasItem(c *gin.Context) {
	vasItemId, err := strconv.Atoi(c.Param("vasItemId"))
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	cart, err := h.cartApp.RemoveVasItem(vasItemId)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Vas item removed successfully",
		"data":    cart,
	})
}
