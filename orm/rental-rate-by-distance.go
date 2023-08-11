package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type RentalRateByDistance struct { // ค่าเช่าตา่มระยะทาง
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	StartDistance  float64 `db:"start_distance" json:"start_distance" gorm:"type:decimal(16,4)"`
	EndDistance    float64 `db:"end_distance" json:"end_distance" gorm:"type:decimal(16,4)"`
	RentalDistance float64 `db:"rental_distance" json:"rental_distance" gorm:"type:decimal(16,4)"`
	VehicleModelID string  `db:"vehicle_model_id" json:"vehicle_model_id" gorm:"type:varchar(36); "`
	Price          float64 `db:"price" json:"price" gorm:"type:decimal(16,4)"`
}
