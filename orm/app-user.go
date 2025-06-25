package orm

import (
	"github.com/FourWD/middleware/model"
)

type AppUser struct {
	ID string `json:"id" query:"id" firestore:"id"`
	model.GormModel
	SlipID string `json:"slip_id" query:"slip_id" firestore:"slip_id"`

	AppUserStatusID  string  `json:"app_user_status_id" query:"app_user_status_id" firestore:"app_user_status_id"`
	AppDriverID      string  `json:"app_driver_id" query:"app_driver_id" firestore:"app_driver_id"`
	DriverName       string  `json:"driver_name" query:"driver_name" firestore:"driver_name"`
	Remark           string  `json:"remark" query:"remark" firestore:"remark"`
	DriverImagePath  string  `json:"driver_image_path" query:"driver_image_path" firestore:"driver_image_path"`
	VehicleImage     string  `json:"vehicle_image" query:"vehicle_image" firestore:"vehicle_image"`
	VehicleColorName string  `json:"vehicle_color_name" query:"vehicle_color_name" firestore:"vehicle_color_name"`
	VehicleCode      string  `json:"vehicle_code" query:"vehicle_code" firestore:"vehicle_code"`
	VehicleModelName string  `json:"vehicle_model_name" firestore:"vehicle_model_name" query:"vehicle_model_name"`
	LicensePlate     string  `json:"license_plate" query:"license_plate" firestore:"license_plate"`
	OriginName       string  `json:"origin_name" query:"origin_name" firestore:"origin_name"`
	DestinationName  string  `json:"destination_name" query:"destination_name" firestore:"destination_name"`
	CustomerID       string  `json:"customer_id" query:"customer_id" firestore:"customer_id"`
	CreditCardNo     string  `json:"credit_card_no" query:"credit_card_no" firestore:"credit_card_no"`
	CreditCardTypeID string  `json:"credit_card_type_id" query:"credit_card_type_id" firestore:"credit_card_type_id"`
	VehiclePrice     float64 `json:"vehicle_price" query:"vehicle_price" firestore:"vehicle_price"`
	IsPaymentFail    bool    `json:"is_payment_fail" firestore:"is_payment_fail" query:"is_payment_fail"`
}
