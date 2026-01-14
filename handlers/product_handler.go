package handlers

import (
	"monolith-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles product-related HTTP requests.
type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler creates a new ProductHandler instance.
func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// GetProducts handles GET requests to retrieve all products.
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct handles GET requests to retrieve a product by ID.
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
