package handlers


import (
    "net/http"
    "strconv"
    "time"

    "nexus/models"
    "nexus/services"

    "github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    var newProduct models.Producto

    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newProduct.ID = services.GenerateID() // Generar un ID aleatorio
    newProduct.FechaCarga = time.Now()
    newProduct.FechaActualizacion = time.Now()

    services.AddProduct(newProduct)

    c.JSON(http.StatusCreated, newProduct)
}

func GetProducts(c *gin.Context) {
    products := services.GetAllProducts()
    c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    product, found := services.GetProductByID(id)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var updatedProduct models.Producto
    if err := c.ShouldBindJSON(&updatedProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedProduct.FechaActualizacion = time.Now()
    if updated := services.UpdateProduct(id, updatedProduct); !updated {
        c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        return
    }

    c.JSON(http.StatusOK, updatedProduct)
}

func DeleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if deleted := services.DeleteProduct(id); !deleted {
        c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}