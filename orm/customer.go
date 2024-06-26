package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type Customer struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); unique;"`
	orm.GormModel

	Code string `json:"code" query:"code" gorm:"type:varchar(20) ; default:null ; index"`

	PrefixID       string    `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"`
	Firstname      string    `json:"firstname" query:"firstname" gorm:"type:varchar(200)"`
	Lastname       string    `json:"lastname" query:"lastname" gorm:"type:varchar(200)"`
	Middlename     string    `json:"middlename" query:"middlename" gorm:"type:varchar(200)"`
	Birthday       time.Time `json:"birthday" query:"birthday" `
	Avatar         string    `json:"avatar" query:"avatar" gorm:"type:varchar(500)"`
	Email          string    `json:"email" query:"email" gorm:"type:varchar(100)"`
	IDCardNo       string    `json:"idcard_no" query:"idcard_no" gorm:"type:varchar(50)"`
	NationalityID  string    `json:"nationality_id" query:"nationality_id" gorm:"type:varchar(3)"` // table country
	PassportNo     string    `json:"passport_no" query:"passport_no" gorm:"type:varchar(20)"`
	RegisterFromID string    `json:"register_from_id" query:"register_from_id" gorm:"type:varchar(2)"` // 01 = สมัครจาก counter สุวรรภูมิ 02 = จากเว็บไซต์ aot
	RunningNo      int       `json:"running_no" query:"running_no" gorm:"primary_key;auto_increment;not_null"`
	PhoneNo1       string    `json:"phone_no_1" query:"phone_no_1" gorm:"type:varchar(20);column:phone_no_1"`
	PhoneNo2       string    `json:"phone_no_2" query:"phone_no_2" gorm:"type:varchar(20);column:phone_no_2"`
	PhoneNo3       string    `json:"phone_no_3" query:"phone_no_3" gorm:"type:varchar(20);column:phone_no_3"`
	IsShow         bool      `json:"is_show" query:"is_show" gorm:"type:bool; comment:'โชว์แสดงผลการค้นหาสมาชิก'"`
}
