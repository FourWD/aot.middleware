package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type CancelType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name string `db:"name" json:"name" gorm:"type:varchar(150)"`
	orm.GormRowOrder
}
