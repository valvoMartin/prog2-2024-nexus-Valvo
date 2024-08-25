package models

import "time"

type Caracteristica struct {
    Nombre string `json:"nombre" binding:"required"`
    Valor  string `json:"valor" binding:"required"`
}

type Producto struct {
    ID                 int              `json:"id"`
    Nombre             string           `json:"nombre" binding:"required"`
    FechaCarga         time.Time        `json:"fechaCarga"`
    FechaActualizacion time.Time        `json:"fechaActualizacion"`
    Caracteristicas    []Caracteristica `json:"caracteristicas" binding:"required,dive,required"`
}