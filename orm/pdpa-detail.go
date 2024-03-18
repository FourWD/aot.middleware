package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type PDPADetail struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	PdpaID string `db:"pdpa_id" json:"pdpa_id" gorm:"type:varchar(36);"`

	Subject string `db:"subject" json:"subject" gorm:"type:varchar(200);"`
	Detail  string `db:"detail" json:"detail" gorm:"type:text;"`
}
