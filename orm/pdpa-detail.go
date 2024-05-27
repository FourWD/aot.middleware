package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type PDPADetail struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	PdpaID string `json:"pdpa_id" query:"pdpa_id" gorm:"type:varchar(36);"`

	Subject string `json:"subject" query:"subject" gorm:"type:varchar(200);"`
	Detail  string `json:"detail" query:"detail" gorm:"type:text;"`
}
