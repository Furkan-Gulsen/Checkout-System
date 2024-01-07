package api

import (
	"strconv"

	"github.com/Furkan-Gulsen/Checkout-System/internal/application"
	"github.com/Furkan-Gulsen/Checkout-System/internal/interfaces/dto"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryApp application.CategoryAppInterface
}

func NewCategoryHandler(categoryApp application.CategoryAppInterface) *CategoryHandler {
	return &CategoryHandler{
		categoryApp: categoryApp,
	}
}

// @Summary List categories
// @Description Get a list of categories
// @Tags Category
// @Produce json
// @Success 200 {object} []entity.Category
// @Router /api/v1/category/list [get]
func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.categoryApp.List()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": categories,
	})
}

// @Summary Create a category
// @Description Create a new category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body dto.CategoryRequest true "Category object"
// @Success 200 {object} string
// @Router /api/v1/category [post]
func (h *CategoryHandler) Create(c *gin.Context) {
	var data dto.CategoryRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	dataEntity := data.ToEntity()
	err := dataEntity.Validate()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	category, createErr := h.categoryApp.Create(&dataEntity)
	if createErr != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Category created successfully", "data": category})
}

// @Summary Get a category by ID
// @Description Get a category by ID
// @Tags Category
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} string
// @Router /api/v1/category/{id} [get]
func (h *CategoryHandler) GetById(c *gin.Context) {
	paramID := c.Param("id")
	if paramID == "" {
		c.JSON(400, gin.H{"message": "Category ID is required"})
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID format"})
		return
	}

	category, err := h.categoryApp.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Category found",
		"data":    category,
	})
}
