package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type Customer struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); unique;"`
	orm.GormModel

	Code string `json:"code" query:"code" gorm:"type:varchar(20) ; default:null ; index"`

	PrefixID        string    `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"`
	Firstname       string    `json:"firstname" query:"firstname" gorm:"type:varchar(200)"`
	Lastname        string    `json:"lastname" query:"lastname" gorm:"type:varchar(200)"`
	Birthday        time.Time `json:"birthday" query:"birthday" `
	Avatar          string    `json:"avatar" query:"avatar" gorm:"type:varchar(500)"`
	Email           string    `json:"email" query:"email" gorm:"type:varchar(50)"`
	IDCardNo        string    `json:"idcard_no" query:"idcard_no" gorm:"type:varchar(50)"`
	NationalityID   string    `json:"nationality_id" query:"nationality_id" gorm:"type:varchar(3)"` // table country
	PassportNo      string    `json:"passport_no" query:"passport_no" gorm:"type:varchar(20)"`
	RegisterFrom    string    `json:"register_from" query:"register_from" gorm:"type:varchar(2)"` // 01 = สมัครจาก counter สุวรรภูมิ 02 = จากเว็บไซต์ aot
	CompanyName     string    `json:"company_name" query:"company_name" gorm:"type:varchar(300)"`
	TaxNo           string    `json:"tax_no" query:"tax_no" gorm:"type:varchar(20)"`
	Address         string    `json:"address" query:"address" gorm:"type:varchar(255)"`
	AddressBranch   string    `json:"address_branch" query:"address_branch" gorm:"type:varchar(255)"`
	AddressCustomer string    `json:"address_customer" query:"address_customer" gorm:"type:varchar(255)"`
	IsShow          bool      `json:"is_show" query:"is_show" gorm:"type:tinyint(1); comment:'โชว์แสดงผลการค้นหาสมาชิก'"`
	IsTax           bool      `json:"is_tax" query:"is_tax" gorm:"type:tinyint(1)"`
	RunningNo       int       `json:"running_no" query:"running_no" gorm:"primary_key;auto_increment;not_null"`
	PhoneNo1        string    `json:"phone_no_1" query:"phone_no_1" gorm:"type:varchar(20);column:phone_no_1"`
	PhoneNo2        string    `json:"phone_no_2" query:"phone_no_2" gorm:"type:varchar(20);column:phone_no_2"`
	PhoneNo3        string    `json:"phone_no_3" query:"phone_no_3" gorm:"type:varchar(20);column:phone_no_3"`
	Postcode        string    `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`
}
