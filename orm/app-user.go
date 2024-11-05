package orm

import (
	"github.com/FourWD/middleware/model"
)

type AppUser struct {
	ID string `json:"id" query:"id" firestore:"id" gorm:"-"`
	model.GormModel
	SlipID string `json:"slip_id" query:"slip_id" firestore:"slip_id" gorm:"-"`

	AppUserStatusID  string  `json:"app_user_status_id" query:"app_user_status_id" firestore:"app_user_status_id" gorm:"-"`
	AppDriverID      string  `json:"app_driver_id" query:"app_driver_id" firestore:"app_driver_id" gorm:"-"`
	DriverName       string  `json:"driver_name" query:"driver_name" firestore:"driver_name" gorm:"-"`
	Remark           string  `json:"remark" query:"remark" firestore:"remark" gorm:"-"`
	DriverImagePath  string  `json:"driver_image_path" query:"driver_image_path" firestore:"driver_image_path" gorm:"-"`
	VehicleImage     string  `json:"vehicle_image" query:"vehicle_image" firestore:"vehicle_image" gorm:"-"`
	VehicleColorName string  `json:"vehicle_color_name" query:"vehicle_color_name" firestore:"vehicle_color_name" gorm:"-"`
	VehicleCode      string  `json:"vehicle_code" query:"vehicle_code" firestore:"vehicle_code" gorm:"-"`
	VehicleModelName string  `json:"vehicle_model_name" firestore:"vehicle_model_name" query:"vehicle_model_name" gorm:"-"`
	LicensePlate     string  `json:"license_plate" query:"license_plate" firestore:"license_plate" gorm:"-"`
	OriginName       string  `json:"origin_name" query:"origin_name" firestore:"origin_name" gorm:"-"`
	DestinationName  string  `json:"destination_name" query:"destination_name" firestore:"destination_name" gorm:"-"`
	CustomerID       string  `json:"customer_id" query:"customer_id" firestore:"customer_id" gorm:"-"`
	VehiclePrice     float64 `json:"vehicle_price" query:"vehicle_price" firestore:"vehicle_price" gorm:"-"`
}
