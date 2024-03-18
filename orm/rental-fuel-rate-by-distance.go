package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type RentalFuelRateByDistance struct { // ค่าเช่าค่าน้ำมันตามระยะทาง
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	StartDistance float64 `db:"start_distance" json:"start_distance" gorm:"type:decimal(16,4)"`
	EndDistance   float64 `db:"end_distance" json:"end_distance" gorm:"type:decimal(16,4)"`

	VehicleModelID string  `db:"vehicle_model_id" json:"vehicle_model_id" gorm:"type:varchar(36); "`
	Litre          float64 `db:"litre" json:"litre" gorm:"type:decimal(16,4)"`
}
