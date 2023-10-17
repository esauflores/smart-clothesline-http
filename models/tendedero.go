package models

type Tendedero struct {
	Id     string `json:"id" gorm:"primaryKey; type:varchar(36)"`
	Nombre string `json:"nombre" gorm:"not null; uniqueIndex: idx_nombre_usuarioid"`
	Modo   bool   `json:"modo" gorm:"not null"`
	Estado bool   `json:"estado" gorm:"not null"`

	UsuarioId string `json:"usuario_id" gorm:"type:varchar(36); uniqueIndex: idx_nombre_usuarioid"`

	Eventos []Evento `json:"eventos" gorm:"foreignKey:TendederoId"`
}
