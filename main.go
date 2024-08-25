package main

import (
    "nexus/handlers"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.POST("/productos", handlers.CreateProduct)
    r.GET("/productos", handlers.GetProducts)
    r.GET("/productos/:id", handlers.GetProductByID)
    r.PUT("/productos/:id", handlers.UpdateProduct)
    r.DELETE("/productos/:id", handlers.DeleteProduct)

    r.Run(":8080")
}