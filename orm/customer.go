package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type Customer struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36); unique;"`
	orm.GormModel

	Code string `db:"code" json:"code" gorm:"type:varchar(20) ; default:null ; index"`

	PrefixID        string    `db:"prefix_id" json:"prefix_id" gorm:"type:varchar(2)"`
	Firstname       string    `db:"firstname" json:"firstname" gorm:"type:varchar(200)"`
	Lastname        string    `db:"lastname" json:"lastname" gorm:"type:varchar(200)"`
	Birthday        time.Time `db:"birthday" json:"birthday" `
	Avatar          string    `db:"avatar" json:"avatar" gorm:"type:varchar(500)"`
	Email           string    `db:"email" json:"email" gorm:"type:varchar(50)"`
	IDCardNo        string    `db:"idcard_no" json:"idcard_no" gorm:"type:varchar(50)"`
	NationalityID   string    `db:"nationality_id" json:"nationality_id" gorm:"type:varchar(3)"` // table country
	PassportNo      string    `db:"passport_no" json:"passport_no" gorm:"type:varchar(20)"`
	RegisterFrom    string    `db:"register_from" json:"register_from" gorm:"type:varchar(2)"` // 01 = สมัครจาก counter สุวรรภูมิ 02 = จากเว็บไซต์ aot
	CompanyName     string    `db:"company_name" json:"company_name" gorm:"type:varchar(300)"`
	TaxNo           string    `db:"tax_no" json:"tax_no" gorm:"type:varchar(20)"`
	Address         string    `db:"address" json:"address" gorm:"type:varchar(255)"`
	AddressBranch   string    `db:"address_branch" json:"address_branch" gorm:"type:varchar(255)"`
	AddressCustomer string    `db:"address_customer" json:"address_customer" gorm:"type:varchar(255)"`
	IsShow          bool      `db:"is_show" json:"is_show" gorm:"type:tinyint(1); comment:'โชว์แสดงผลการค้นหาสมาชิก'"`
	IsTax           bool      `db:"is_tax" json:"is_tax" gorm:"type:tinyint(1)"`
	RunningNo       int       `db:"running_no" json:"running_no" gorm:"primary_key;auto_increment;not_null"`
	PhoneNo1        string    `db:"phone_no_1" json:"phone_no_1" gorm:"type:varchar(20);column:phone_no_1"`
	PhoneNo2        string    `db:"phone_no_2" json:"phone_no_2" gorm:"type:varchar(20);column:phone_no_2"`
	PhoneNo3        string    `db:"phone_no_3" json:"phone_no_3" gorm:"type:varchar(20);column:phone_no_3"`
	Postcode        string    `db:"postcode" json:"postcode" gorm:"type:varchar(5)"`
}
