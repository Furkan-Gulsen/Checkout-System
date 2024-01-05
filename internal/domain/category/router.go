package category

import (
	"github.com/Furkan-Gulsen/Checkout-System/internal/entities"
	"github.com/Furkan-Gulsen/Checkout-System/internal/infra/database"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup, db *database.Database, dbName string) {
	r := g.Group("/category")

	categoryRepository := NewCategoryRepository(db, dbName)
	categoryService := NewCategoryService(*categoryRepository)
	categoryRouter := NewCategoryRouter(*categoryService)

	r.GET("/list", categoryRouter.List)
	r.POST("/create", categoryRouter.Create)
}

type CategoryRouter struct {
	categoryService CategoryService
}

func NewCategoryRouter(categoryService CategoryService) *CategoryRouter {
	return &CategoryRouter{
		categoryService: categoryService,
	}
}

// @Summary List categories
// @Description Get a list of categories
// @Tags Category
// @Produce json
// @Success 200 {object} []entities.Category
// @Router /api/v1/category/list [get]
func (r *CategoryRouter) List(c *gin.Context) {
	categories, err := r.categoryService.List()
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
// @Param category body entities.Category true "Category object"
// @Success 200 {string} string "Category created successfully"
// @Router /api/v1/category/create [post]
func (r *CategoryRouter) Create(c *gin.Context) {
	var category entities.Category

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

	if err := r.categoryService.Create(category); err != nil {
		c.JSON(500, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": true, "message": "Category created successfully"})
}
