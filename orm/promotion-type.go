package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type PromotionType struct { //no CRUD 01 = เงินสด 02 = creditcard 03= chuqe 04 = banktransfer
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `db:"name" json:"name" gorm:"type:varchar(150)"`
	orm.GormRowOrder
}
