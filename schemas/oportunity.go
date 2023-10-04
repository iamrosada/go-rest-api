package schemas

import (
	"gorm.io/gorm"
)

type Oportunity struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
