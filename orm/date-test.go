package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type DateTest struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DateAt string `db:"date_at" json:"date_at" gorm:"default:null;  column:date_at;  type:datetime; comment:'วันเวลาที่ยกเลิก' "`
}
