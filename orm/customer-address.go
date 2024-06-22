package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type CustomerAddress struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	CustomerID   string `json:"customer_id"  query:"customer_id" gorm:"type:varchar(36); "`
	CompanyName  string `json:"company_name" query:"company_name" gorm:"type:varchar(255)"`
	BranchName   string `json:"branch_name" query:"branch_name" gorm:"type:varchar(255)"`
	TaxNo        string `json:"tax_no" query:"tax_no" gorm:"type:varchar(50)"`
	IsHQ         bool   `json:"is_hq" query:"is_hq" gorm:"type:bool"` // is_hq = 1 เป็นสำนักงาน
	Address      string `json:"address" query:"address" gorm:"type:text"`
	Postcode     string `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`
	CompanyPhone string `json:"company_phone" query:"company_phone" gorm:"type:varchar(50)"`
}
