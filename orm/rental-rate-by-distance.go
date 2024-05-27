package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type RentalRateByDistance struct { // ค่าเช่าตา่มระยะทาง
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	StartDistance float64 `json:"start_distance" query:"start_distance" gorm:"type:decimal(16,4)"`
	EndDistance   float64 `json:"end_distance" query:"end_distance" gorm:"type:decimal(16,4)"`

	RentalDistance float64 `json:"rental_distance" query:"rental_distance" gorm:"type:decimal(16,4)"`
	VehicleModelID string  `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36); "`
	Price          float64 `json:"price" query:"price" gorm:"type:decimal(16,4)"`
}
