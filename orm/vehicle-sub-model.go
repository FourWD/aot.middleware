package orm

import orm "github.com/FourWD/middleware/orm"

type VehicleSubModel struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	VehicleSourceID string `db:"vehicle_source_id" json:"vehicle_source_id" gorm:"type:varchar(36)"`
	VehicleModelID  string `db:"vehicle_model_id" json:"vehicle_model_id" gorm:"type:varchar(36)"`

	Name     string `db:"name" json:"type_name"`
	NameEn   string `db:"name_en" json:"type_name_en"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
	Image    string `db:"image" json:"image"`
	//
}
