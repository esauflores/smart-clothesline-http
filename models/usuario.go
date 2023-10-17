package models

type Usuario struct {
	Id         string `json:"id" gorm:"primaryKey; type:varchar(36)"`
	Nombres    string `json:"nombres" gorm:"not null; type:varchar(50)"`
	Apellidos  string `json:"apellidos" gorm:"not null; type:varchar(50)"`
	Email      string `json:"email" gorm:"uniqueIndex; not null; type:varchar(255)"`
	Contrasena string `json:"contrasena" gorm:"not null; type:varchar(255)"`

	Tendederos []Tendedero `json:"tendederos" gorm:"foreignKey:UsuarioId"`
}
