package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type RentalFuelRateByDistance struct { // ค่าเช่าค่าน้ำมันตามระยะทาง
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	StartDistance float64 `json:"start_distance" query:"start_distance" gorm:"type:decimal(16,4)"`
	EndDistance   float64 `json:"end_distance" query:"end_distance" gorm:"type:decimal(16,4)"`

	VehicleModelID string  `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36); "`
	Litre          float64 `json:"litre" query:"litre" gorm:"type:decimal(16,4)"`
}
