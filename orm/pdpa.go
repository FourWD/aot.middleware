package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type PDPA struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name    string `db:"name" json:"name" gorm:"type:varchar(150)"`
	Version string `db:"version" json:"version" gorm:"type:varchar(10)"`
}
