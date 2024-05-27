package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Running struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	RunningNo   string `json:"running_no" query:"running_no" gorm:"type:varchar(6)"`
	TempCommuID string `json:"tempcommu_id" query:"tempcommu_id" gorm:"type:varchar(36);"`
	TempMuID    string `json:"tempmu_id" query:"tempmu_id" gorm:"type:varchar(36);"`
}
