package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type SlipDriverCancelType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(150)"`
	orm.GormRowOrder
}
