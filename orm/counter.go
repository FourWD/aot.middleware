package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Counter struct { // จุดขายหน้าเค้าเตอร์  ที่สุวรรณภูมิ
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(150)"`
	IP   string `json:"ip" query:"ip" gorm:"type:varchar(20)"`
}
