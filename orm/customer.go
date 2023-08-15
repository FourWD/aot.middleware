package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type Customer struct {
	ID string `json:"id" gorm:"type:varchar(36); unique;"`
	orm.GormModel

	Code string `json:"code" gorm:"type:varchar(20) ; dafault:null ; index"`

	PrefixID        string `json:"prefix_id" gorm:"type:varchar(2)"`
	Firstname       string `json:"firstname" gorm:"type:varchar(50)"`
	Lastname        string `json:"lastname" gorm:"type:varchar(50)"`
	Birthday        string `json:"birthday" `
	Avatar          string `json:"avatar" gorm:"type:varchar(100)"`
	Email           string `json:"email" gorm:"type:varchar(50)"`
	IDCardNo        string `json:"idcard_no" gorm:"type:varchar(13)"`
	NationalityID   string `json:"nationality_id" gorm:"type:varchar(3)"` // table country
	PassportNo      string `json:"passport_no" gorm:"type:varchar(13)"`
	RegisterFrom    string `json:"register_from" gorm:"type:varchar(2)"` // 01 = สมัครจาก counter สุวรรภูมิ 02 = จากเว็บไซต์ aot
	CompanyName     string `json:"company_name" gorm:"type:varchar(150)"`
	TaxNo           string `json:"tax_no" gorm:"type:varchar(20)"`
	Address         string `json:"address" gorm:"type:varchar(255)"`
	AddressBranch   string `json:"address_branch" gorm:"type:varchar(255)"`
	AddressCustomer string `json:"address_customer" gorm:"type:varchar(255)"`
	IsShow          int    `json:"is_show" gorm:"type:tinyint(2);  comment:'โชว์แสดงผลการค้นหาสมาชิก'"`
	IsTax           int    `json:"is_tax" gorm:"type:tinyint(2); comment:'ใบกำกับภาษี none company'"` //is_hq
	RunningNo       int    `json:"running_no"  gorm:"primary_key;auto_increment;not_null"`
	PhoneNo1        string `json:"phone_no1" gorm:"type:varchar(10)"`
	PhoneNo2        string `json:"phone_no2" gorm:"type:varchar(10)"`
	Postcode        string `json:"postcode" gorm:"type:varchar(5)"`
}
