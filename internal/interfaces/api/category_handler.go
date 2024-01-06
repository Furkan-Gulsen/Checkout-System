package api

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/Furkan-Gulsen/Checkout-System/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
	categoryApp application.CategoryAppInterface
}

func NewCategoryRouter(categoryApp application.CategoryAppInterface) *CategoryRouter {
	return &CategoryRouter{
		categoryApp: categoryApp,
	}
}

// @Summary List categories
// @Description Get a list of categories
// @Tags Category
// @Produce json
// @Success 200 {object} []entity.Category
// @Router /api/v1/category/list [get]
func (h *CategoryRouter) List(c *gin.Context) {
	categories, err := h.categoryApp.List()
	if err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": categories,
	})
}

// @Summary Create a category
// @Description Create a new category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body entity.Category true "Category object"
// @Success 200 {string} string "Category created successfully"
// @Router /api/v1/category/create [post]
func (h *CategoryRouter) Create(c *gin.Context) {
	var category entity.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}

	err := category.Validate()
	if err != nil {
		c.JSON(400, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if err := h.categoryApp.Create(category); err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "Category created successfully"})
}
