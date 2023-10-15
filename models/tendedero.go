package models

type Tendedero struct {
	Id     string `json:"id" gorm:"primaryKey; type:varchar(36)"`
	Modo   bool   `json:"modo" gorm:"not null"`
	Estado bool   `json:"estado" gorm:"not null"`

	// foreign key
	UsuarioId string `json:"usuario_id" gorm:"type:varchar(36)"`
}
