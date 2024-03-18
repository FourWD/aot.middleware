package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type FuelBill struct { //no CRUD
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	SourceID string `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`

	RefuelAt            string  `db:"refuel_at"  json:"refuel_at" gorm:"default:null; type:varchar(50); comment:'วันที่เติมน้ำมัน' "`
	VehicleID           string  `db:"vehicle_id" json:"vehicle_id" gorm:"default:null; type:varchar(36); "`
	DriverID            string  `db:"driver_id" json:"driver_id" gorm:"default:null; type:varchar(36); "`
	GasStationID        string  `db:"gas_station_id" json:"gas_station_id" gorm:"default:null; type:varchar(36); "`
	ForceGasStationName string  `db:"force_gas_station_name" json:"force_gas_station_name" gorm:"default:null; type:varchar(200); "`
	FuelID              string  `db:"fuel_id" json:"fuel_id" gorm:"type:varchar(2);"`
	Litre               float64 `db:"litre" json:"litre" gorm:"default:null; type:decimal(16,4); comment:'จำนวนน้ำมันที่เติม' "`
	Price               float64 `db:"price" json:"price" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าน้ำมัน' "`
}
