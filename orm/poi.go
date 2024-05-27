package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type Poi struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name       string  `json:"name" query:"name" gorm:"type:varchar(200)"`
	NameEn     string  `json:"name_en" query:"name_en" gorm:"type:varchar(200)"`
	Address    string  `json:"address" query:"address" gorm:"type:varchar(200)"`
	AddressEn  string  `json:"address_en" query:"address_en" gorm:"type:varchar(200)"`
	PhoneNo    string  `json:"phone_no" query:"position" gorm:"type:varchar(10)"`
	Distance   float64 `json:"distance" query:"distance" gorm:"type:decimal(16,4)"`
	Duration   int     `json:"duration" query:"duration" gorm:"type:int"`
	ProvinceID string  `json:"province_id" query:"province_id" gorm:"type:varchar(2)"`
	PoiTypeID  string  `json:"poi_type_id" query:"poi_type_id" gorm:"type:varchar(2)"`
	IsActive   bool    `json:"is_active" query:"is_active" gorm:"type:tinyint(2)"`
	Latitude   float64 `json:"lat" query:"lat" gorm:"type:decimal(10,6)"`
	Longitude  float64 `json:"long" query:"long" gorm:"type:decimal(10,6)"`
	Image1     string  `json:"image1"  query:"image1" gorm:"default:null; type:varchar(200)"`
	Image2     string  `json:"image2"  query:"image2" gorm:"default:null; type:varchar(200)"`
	Image3     string  `json:"image3"  query:"image3" gorm:"default:null; type:varchar(200)"`
	Image4     string  `json:"image4"  query:"image4" gorm:"default:null; type:varchar(200)"`
	Image5     string  `json:"image5"  query:"image5" gorm:"default:null; type:varchar(200)"`
}
