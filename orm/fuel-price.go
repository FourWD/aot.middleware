package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type FuelPrice struct { //no CRUD
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	FuelID    string  `db:"fuel_id" json:"fuel_id" gorm:"type:varchar(2);"`
	StartDate string  `db:"start_date" json:"start_date" gorm:"type:varchar(50) ; comment:'วันที่เริ่ม' "`
	EndDate   string  `db:"end_date" json:"end_date" gorm:"type:varchar(50) ; comment:'วันที่จบ'"`
	Month     int     `db:"imonth" json:"imonth" gorm:"type:int(2);"`
	Year      int     `db:"iyear" json:"iyear" gorm:"type:int(4);"`
	Average   float64 `db:"average" json:"average" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าน้ำมันแต่ละเดือน' "`
}
