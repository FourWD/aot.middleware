package orm

type AppDriver struct {
	Slip
	// AppDriverStatusID string `json:"app_driver_status_id" query:"app_driver_status_id" firestore:"app_driver_status_id" gorm:"type:varchar(2);"`
	VehicleSubTypeID string `json:"vehicle_sub_type_id" firestore:"vehicle_sub_type_id" gorm:"-"`
	VehicleModelID   string `json:"vehicle_model_id" firestore:"vehicle_model_id" gorm:"-"`
	SlipID           string `json:"slip_id" firestore:"slip_id" query:"slip_id" gorm:"-"`
	VehicleName      string `json:"vehicle_name" firestore:"vehicle_name" query:"vehicle_name" gorm:"-"`
	IsOriginAirport  bool   `json:"is_origin_airport" firestore:"is_origin_airport" query:"is_origin_airport" gorm:"-"`
	// Name                string `json:"name"  firestore:"name" `
	// VehicleID           string `json:"vehicle_id"  firestore:"driver_register_id" `
	// VehicleCode         string `json:"vehicle_code"  firestore:"vehicle_id"`
	// VehicleColorName    string `json:"vehicle_color_name"  firestore:"vehicle_color_name"`
	// VehicleName         string `json:"vehicle_name"  firestore:"vehicle_name" `
	// VehicleTypeName     string `json:"vehicle_type_name"  firestore:"vehicle_type_name"`
	// VehicleSubModelName string `json:"vehicle_sub_model_name"  firestore:"vehicle_sub_model_name"`
}
