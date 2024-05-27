package orm

import orm "github.com/FourWD/middleware/model"

type VehicleSubModel struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	VehicleSourceID string `json:"vehicle_source_id" query:"vehicle_source_id" gorm:"type:varchar(36)"`
	VehicleModelID  string `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	Name            string `json:"name" query:"type_name"`
	NameEn          string `json:"name_en" query:"type_name_en"`
	Image           string `json:"image" query:"image"`
	IsShow          bool   `json:"is_show" query:"is_show" gorm:"type:bool"`
	SubModelGroup   string `json:"sub_model_group" query:"sub_model_group" gorm:"type:varchar(5)"`
	orm.GormRowOrder
}
