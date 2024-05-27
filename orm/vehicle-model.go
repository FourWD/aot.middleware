package orm

import orm "github.com/FourWD/middleware/model"

type VehicleModel struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	VehicleBrandID string `json:"vehicle_brand_id" query:"vehicle_brand_id"`
	VehicleTypeID  string `json:"vehicle_type_id" query:"vehicle_type_id"`
	Name           string `json:"name" query:"type_name"`
	NameEn         string `json:"name_en" query:"type_name_en"`
	Image          string `json:"image" query:"image"`
	orm.GormRowOrder
}
