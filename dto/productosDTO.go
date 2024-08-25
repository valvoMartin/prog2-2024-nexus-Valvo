package dto


import "nexus/models"

type CreateProductDTO struct {
    Nombre          string          `json:"nombre" binding:"required"`
    Caracteristicas []models.Caracteristica `json:"caracteristicas" binding:"required,dive,min=1"`
}


type UpdateProductDTO struct {
    Nombre          *string          `json:"nombre"`
    Caracteristicas *[]models.Caracteristica `json:"caracteristicas"`
}