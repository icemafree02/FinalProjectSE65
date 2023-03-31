package orm

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Carname string
	Detail  string
	Image   string
}
