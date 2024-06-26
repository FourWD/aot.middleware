package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type SlipSubType struct { //ประเภทของสลิป 01 = oneway 02 = roundtrip 03 = byhours
	ID string `query:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name string `query:"name" json:"name" gorm:"type:varchar(150)"`
	orm.GormRowOrder
}
