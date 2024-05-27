package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Configs struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name        string `json:"name" query:"name" gorm:"type:varchar(150)"`
	ConfigValue string `json:"config_value" query:"config_value" gorm:"type:varchar(150)"`
} // void ผ่าน callcenter 100 % // void ผ่าน website 95%
