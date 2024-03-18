package orm

import orm "github.com/FourWD/middleware/model"

type VehicleModel struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	VehicleBrandID string `db:"vehicle_brand_id" json:"vehicle_brand_id"`
	VehicleTypeID  string `db:"vehicle_type_id" json:"vehicle_type_id"`
	Name           string `db:"name" json:"type_name"`
	NameEn         string `db:"name_en" json:"type_name_en"`
	Image          string `db:"image" json:"image"`
	RowOrder       int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
