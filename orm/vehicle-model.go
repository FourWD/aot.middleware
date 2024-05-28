package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type VehicleModel struct {
	orm.VehicleModel
	Image string `json:"image" query:"image" gorm:"type:varchar(200);"`
}
