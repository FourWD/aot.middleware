package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Vehicle struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	VehicleSubmodelID string  `json:"vehicle_submodel_id" query:"vehicle_submodel_id" gorm:"type:varchar(2)"`
	Code              string  `json:"code" query:"code" gorm:"type:varchar(50)"`
	LicensePlate      string  `json:"license_plate" query:"license_plate" gorm:"type:varchar(50)"`
	Latitude          float64 `json:"lat" query:"lat" gorm:"type:decimal(10,6)"`
	Longitude         float64 `json:"long" query:"long" gorm:"type:decimal(10,6)"`
	VehicleStatusID   string  `json:"vehicle_status_id" query:"vehicle_status_id" gorm:"type:varchar(2);"`
	LocationTypeID    string  `json:"location_type_id" query:"location_type_id" gorm:"type:varchar(2);"`
	CurrentDriverID   string  `json:"current_driver_id" query:"current_driver_id" gorm:"default:null; type:varchar(36); "`
}
