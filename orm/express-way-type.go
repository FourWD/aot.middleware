package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type ExpresswayType struct { //no CRUD ทางด่วน 2 ประเภท
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel
	Name string `json:"name" query:"name" gorm:"type:varchar(50)"`
	orm.GormRowOrder
}
