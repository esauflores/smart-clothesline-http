package models

type Tendedero struct {
	Id     string `json:"id" gorm:"primaryKey"`
	Modo   bool   `json:"modo" gorm:"not null" binding:"required,notnull"`
	Estado bool   `json:"estado" gorm:"not null" binding:"required,notnull"`
}
