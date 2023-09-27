package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Capture struct {
	gorm.Model
	Id              string `gorm:"primary_key"`
	TipoImovel      string
	UsoImovel       string
	Imovel          string
	Caracteristicas string
	CreatedAt       time.Time
}
