// Package main is the entry point for the monolith application server.
package main

import (
	"log"
	"monolith-app/handlers"
	"monolith-app/services"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	userService := services.NewUserService()
	productService := services.NewProductService()
	orderService := services.NewOrderService(userService, productService)

	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)
	healthHandler := handlers.NewHealthHandler()

	api := r.Group("/api/v1")
	{
		api.GET("/health", healthHandler.Health)

		api.GET("/users", userHandler.GetUsers)
		api.GET("/users/:id", userHandler.GetUser)

		api.GET("/products", productHandler.GetProducts)
		api.GET("/products/:id", productHandler.GetProduct)

		api.POST("/orders", orderHandler.CreateOrder)
		api.GET("/orders", orderHandler.GetOrders)
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Println("Monolith application starting on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
