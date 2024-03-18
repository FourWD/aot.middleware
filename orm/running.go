package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Running struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	RunningNo   string `db:"running_no" json:"running_no" gorm:"type:varchar(6)"`
	TempCommuID string `db:"tempcommu_id" json:"tempcommu_id" gorm:"type:varchar(36);"`
	TempMuID    string `db:"tempmu_id" json:"tempmu_id" gorm:"type:varchar(36);"`
}
