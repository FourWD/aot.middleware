package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type JobStatus struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(10);primary_key;"`
	orm.GormModel

	Name string `db:"name" json:"name" gorm:"type:varchar(20)"`
}
