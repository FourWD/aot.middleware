package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type UserGroup struct { //ประเภทของสลิป 01 = oneway 02 = roundtrip 03 = byhours
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(150)"`
	orm.GormRowOrder
}
