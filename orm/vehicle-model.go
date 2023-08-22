package orm

import orm "github.com/FourWD/middleware/orm"

type VehicleModel struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	VehicleBrandID string `db:"vehicle_brand_id" json:"vehicle_brand_id"`
	VehicleTypeID  string `db:"vehicle_type_id" json:"vehicle_type_id"`
	Name           string `db:"name" json:"type_name"`
	NameEn         string `db:"name_en" json:"type_name_en"`
	Image          string `db:"image" json:"image"`
	IsShow         bool   `db:"is_show" json:"is_show"  gorm:"type:bool"`
	SubModelGroup  string `db:"sub_model_group" json:"sub_model_group" gorm:"type:varchar(5)"`
	RowOrder       int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
