package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type PaymentType struct { //no CRUD 01 = เงินสด 02 = creditcard 03= chuqe 04 = banktransfer
	orm.PaymentType
	IsWht bool `json:"is_wht" query:"is_wht" gorm:"default:0; type:tinyint(1)"`
}
