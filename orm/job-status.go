package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type JobStatus struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(10);primary_key;"`
	orm.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(20)"`
}
