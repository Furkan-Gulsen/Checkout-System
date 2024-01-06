package api

import (
	"strconv"

	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
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
// @Router /api/v1/cart/{cartId}/promotion/{promotionId} [post]
func (h *CartHandler) ApplyPromotion(c *gin.Context) {
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	promotionId, err := strconv.Atoi(c.Param("promotionId"))
	if err != nil {
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = h.cartApp.ApplyPromotion(cartId, promotionId)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Promotion applied successfully",
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
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	cart, err := h.cartApp.Display(cartId)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": cart,
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
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	err = h.cartApp.ResetCart(cartId)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Cart reset successfully",
	})
}
