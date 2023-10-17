package models

type Evento struct {
	Id     string `json:"id" gorm:"primaryKey; type:varchar(36)"`
	Modo   bool   `json:"modo" gorm:"not null"`
	Estado bool   `json:"estado" gorm:"not null"`

	TendederoId string `json:"tendedero_id" gorm:"not null; type:varchar(36)"`
}
