package models

type Datos struct {
	Id       int    `gorm:"primaryKey"`
	Puerta   string `gorm:"type:varchar(11)" json:"puerta"`
	Luz      string `gorm:"type:varchar(11)" json:"luz"`
	Personas string `gorm:"type:varchar(11)" json:"personas"`
}
