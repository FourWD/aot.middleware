package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type FuelBill struct { //no CRUD
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	SourceID string `json:"source_id" query:"source_id" gorm:"type:varchar(2);"`

	RefuelAt            time.Time `json:"refuel_at"  query:"refuel_at" gorm:"default:null;comment:'วันที่เติมน้ำมัน'"`
	VehicleID           string    `json:"vehicle_id" query:"vehicle_id" gorm:"default:null;type:varchar(36);"`
	DriverID            string    `json:"driver_id" query:"driver_id" gorm:"default:null;type:varchar(36);"`
	GasStationID        string    `json:"gas_station_id" query:"gas_station_id" gorm:"default:null;type:varchar(36); "`
	ForceGasStationName string    `json:"force_gas_station_name" query:"force_gas_station_name" gorm:"default:null;type:varchar(200); "`
	FuelID              string    `json:"fuel_id" query:"fuel_id" gorm:"type:varchar(2);"`
	Litre               float64   `json:"litre" query:"litre" gorm:"default:null; type:decimal(16,4); comment:'จำนวนน้ำมันที่เติม'"`
	Price               float64   `json:"price" query:"price" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าน้ำมัน'"`
}
