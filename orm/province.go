package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type Province struct {
	orm.Province
	ProvinceTypeID string `json:"province_type_id" query:"province_type_id" gorm:"type:varchar(2);"`
}
