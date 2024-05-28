package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Slip struct {
	SlipModel

	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	JobID string `json:"job_id" query:"job_id" gorm:"type:varchar(10);comment:'หมายเลขงาน'"` //หมายเลขงาน
}
