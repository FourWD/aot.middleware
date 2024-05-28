package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type VehicleType struct {
	orm.VehicleType
	Image string `json:"image" query:"image" gorm:"default:null; type:varchar(200);"`
}
