package services

import (
    "math/rand"
    "nexus/models"
    "sync"
    "time"
)

var (
    products = make(map[int]models.Producto)
    mu       sync.Mutex
)

func GenerateID() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(1000000)
}

func AddProduct(product models.Producto) {
    mu.Lock()
    defer mu.Unlock()
    products[product.ID] = product
}

func GetAllProducts() []models.Producto {
    mu.Lock()
    defer mu.Unlock()
    productList := make([]models.Producto, 0, len(products))
    for _, product := range products {
        productList = append(productList, product)
    }
    return productList
}

func GetProductByID(id int) (models.Producto, bool) {
    mu.Lock()
    defer mu.Unlock()
    product, found := products[id]
    return product, found
}

func UpdateProduct(id int, updatedProduct models.Producto) bool {
    mu.Lock()
    defer mu.Unlock()
    if _, found := products[id]; !found {
        return false
    }
    updatedProduct.ID = id
    products[id] = updatedProduct
    return true
}

func DeleteProduct(id int) bool {
    mu.Lock()
    defer mu.Unlock()
    if _, found := products[id]; !found {
        return false
    }
    delete(products, id)
    return true
}