package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Poi struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name       string  `db:"name" json:"name" gorm:"type:varchar(200)"`
	NameEn     string  `db:"name_en" json:"name_en" gorm:"type:varchar(200)"`
	Address    string  `db:"address" json:"address" gorm:"type:varchar(200)"`
	AddressEn  string  `db:"address_en" json:"address_en" gorm:"type:varchar(200)"`
	PhoneNo    string  `db:"phone_no" json:"position" gorm:"type:varchar(10)"`
	Distance   float64 `db:"distance" json:"distance" gorm:"type:decimal(16,4)"`
	Duration   int     `db:"duration" json:"duration" gorm:"type:int"`
	ProvinceID string  `db:"province_id" json:"province_id" gorm:"type:varchar(2)"`
	PoiTypeID  string  `db:"poi_type_id" json:"poi_type_id" gorm:"type:varchar(2)"`
	IsActive   bool    `db:"is_active" json:"is_active" gorm:"type:tinyint(2)"`
	Latitude   float64 `db:"lat" json:"lat" gorm:"type:decimal(10,6)"`
	Longitude  float64 `db:"long" json:"long" gorm:"type:decimal(10,6)"`
	Image1     string  `db:"image1"  json:"image1" gorm:"default:null; type:varchar(200)"`
	Image2     string  `db:"image2"  json:"image2" gorm:"default:null; type:varchar(200)"`
	Image3     string  `db:"image3"  json:"image3" gorm:"default:null; type:varchar(200)"`
	Image4     string  `db:"image4"  json:"image4" gorm:"default:null; type:varchar(200)"`
	Image5     string  `db:"image5"  json:"image5" gorm:"default:null; type:varchar(200)"`
}
