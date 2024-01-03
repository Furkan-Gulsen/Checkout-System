package product

import "github.com/gin-gonic/gin"

func SetupRouter(g *gin.Engine) {
	productRouterGroup := g.Group("/product")

	productRepository := NewProductRepository()
	productService := NewProductService(*productRepository)
	productRouter := NewProductRouter(*productService)

	productRouterGroup.GET("/list", productRouter.list)
	productRouterGroup.POST("/create", productRouter.create)
	productRouterGroup.POST("/update", productRouter.update)
	productRouterGroup.POST("/delete", productRouter.delete)
}

type ProductRouter struct {
	ProductService ProductService
}

func NewProductRouter(productService ProductService) *ProductRouter {
	return &ProductRouter{
		ProductService: productService,
	}
}

func (p *ProductRouter) list(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (p *ProductRouter) create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (p *ProductRouter) update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (p *ProductRouter) delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
