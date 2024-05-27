package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Configs struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name        string `db:"name" json:"name" gorm:"type:varchar(150)"`
	ConfigValue string `db:"config_value" json:"config_value" gorm:"type:varchar(150)"`
} // void ผ่าน callcenter 100 % // void ผ่าน website 95%
