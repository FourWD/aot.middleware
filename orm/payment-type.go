package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type PaymentType struct { //no CRUD 01 = เงินสด 02 = creditcard 03= chuqe 04 = banktransfer
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name  string `json:"name" query:"name" gorm:"type:varchar(150)"`
	IsWht bool   `json:"is_wht" query:"is_wht" gorm:"default:0; type:tinyint(1)"`
	orm.GormRowOrder
}
