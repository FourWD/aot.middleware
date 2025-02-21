package orm

type AppDriver struct {
	Slip
	LoginDriverName         string `json:"login_driver_name" firestore:"login_driver_name" query:"login_driver_name"`
	LoginDriverImage        string `json:"login_driver_image" firestore:"login_driver_image" query:"login_driver_image"`
	LoginVehicleTypeID      string `json:"login_vehicle_type_id" firestore:"login_vehicle_type_id" query:"login_vehicle_type_id"`
	LoginVehicleTypeName    string `json:"login_vehicle_type_name" firestore:"login_vehicle_type_name" query:"login_vehicle_type_name"`
	LoginVehicleSubTypeID   string `json:"login_vehicle_sub_type_id" firestore:"login_vehicle_sub_type_id" query:"login_vehicle_sub_type_id"`
	LoginVehicleSubTypeName string `json:"login_vehicle_sub_type_name" firestore:"login_vehicle_sub_type_name" query:"login_vehicle_sub_type_name"`
	LoginVehicleModelID     string `json:"login_vehicle_model_id" firestore:"login_vehicle_model_id" query:"login_vehicle_model_id"`
	LoginVehicleModelName   string `json:"login_vehicle_model_name" firestore:"login_vehicle_model_name" query:"login_vehicle_model_name"`

	SlipID          string `json:"slip_id" firestore:"slip_id" query:"slip_id"`
	IsOriginAirport bool   `json:"is_origin_airport" firestore:"is_origin_airport" query:"is_origin_airport"`

	LicensePlate       string `json:"license_plate" firestore:"license_plate" query:"license_plate"`
	VehicleCode        string `json:"vehicle_code" firestore:"vehicle_code" query:"vehicle_code"`
	VehicleImage       string `json:"vehicle_image" firestore:"vehicle_image" query:"vehicle_image"`
	VehicleColorName   string `json:"vehicle_color_name" firestore:"vehicle_color_name" query:"vehicle_color_name"`
	VehicleSubTypeID   string `json:"vehicle_sub_type_id" firestore:"vehicle_sub_type_id" query:"vehicle_sub_type_id"`
	VehicleSubTypeName string `json:"vehicle_sub_type_name" firestore:"vehicle_sub_type_name" query:"vehicle_sub_type_name"`
	VehicleModelID     string `json:"vehicle_model_id" firestore:"vehicle_model_id" query:"vehicle_model_id"`
	VehicleModelName   string `json:"vehicle_model_name" firestore:"vehicle_model_name" query:"vehicle_model_name"`

	DriverName  string `json:"driver_name" firestore:"driver_name" query:"driver_name"`
	DriverImage string `json:"driver_image" firestore:"driver_image" query:"driver_image"`
	JwtToken    string `json:"jwt_token" firestore:"jwt_token" query:"jwt_token"`
}
