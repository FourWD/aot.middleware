package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type ExpressWayBill struct { //no CRUD
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel
	SourceID     string    `json:"source_id" query:"source_id" gorm:"type:varchar(2);"`
	UseAt        time.Time `json:"use_at" query:"use_at" gorm:"default:null; comment:'วันที่ขึ้นทางด่วน' "`
	VehicleID    string    `json:"vehicle_id" query:"vehicle_id" gorm:"default:null; type:varchar(36); "`
	DriverID     string    `json:"driver_id" query:"driver_id" gorm:"default:null; type:varchar(36); "`
	ExpresswayID string    `json:"expressway_id" query:"expressway_id" gorm:"type:varchar(2)"`
	IsReturnTrip bool      `json:"is_return_trip" query:"is_return_trip" gorm:"type:tinyint(1);comment:'0 = ขาไป 1 = ขากลับ'"`
	SlipID       string    `json:"slip_id" query:"slip_id" gorm:"type:varchar(36);"`
	Price        float64   `json:"price" query:"price" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าทางด่วน' "`
}
