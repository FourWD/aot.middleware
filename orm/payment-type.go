package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type PaymentType struct { //no CRUD 01 = เงินสด 02 = creditcard 03= chuqe 04 = banktransfer
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name" gorm:"type:varchar(150)"`
	IsWht    bool   `db:"is_wht" json:"is_wht" gorm:"default:0; type:tinyint(1); comment:'มีwht?' "`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
