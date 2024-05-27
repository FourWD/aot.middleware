package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type FuelPrice struct { //no CRUD
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	FuelID    string    `json:"fuel_id" query:"fuel_id" gorm:"type:varchar(2);"`
	StartDate time.Time `json:"start_date" query:"start_date"`
	EndDate   time.Time `json:"end_date" query:"end_date"`
	Month     int       `json:"imonth" query:"imonth" gorm:"type:int(2);"`
	Year      int       `json:"iyear" query:"iyear" gorm:"type:int(4);"`
	Average   float64   `json:"average" query:"average" gorm:"default:null;type:decimal(16,4);comment:'ราคาค่าน้ำมันแต่ละเดือน'"`
}
