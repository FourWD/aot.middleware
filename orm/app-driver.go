package orm

type AppDriver struct {
	Slip
	LoginDriverName         string `json:"login_driver_name" firestore:"login_driver_name" query:"login_driver_name" gorm:"-"`
	LoginVehicleTypeID      string `json:"login_vehicle_type_id" firestore:"login_vehicle_type_id" query:"login_vehicle_type_id" gorm:"-"`
	LoginVehicleTypeName    string `json:"login_vehicle_type_name" firestore:"login_vehicle_type_name" query:"login_vehicle_type_name" gorm:"-"`
	LoginVehicleSubTypeID   string `json:"login_vehicle_sub_type_id" firestore:"login_vehicle_sub_type_id" query:"login_vehicle_sub_type_id" gorm:"-"`
	LoginVehicleSubTypeName string `json:"login_vehicle_sub_type_name" firestore:"login_vehicle_sub_type_name" query:"login_vehicle_sub_type_name" gorm:"-"`
	LoginVehicleModelID     string `json:"login_vehicle_model_id" firestore:"login_vehicle_model_id" query:"login_vehicle_model_id" gorm:"-"`
	LoginVehicleModelName   string `json:"login_vehicle_model_name" firestore:"login_vehicle_model_name" query:"login_vehicle_model_name" gorm:"-"`

	SlipID             string `json:"slip_id" firestore:"slip_id" query:"slip_id" gorm:"-"`
	VehicleSubTypeID   string `json:"vehicle_sub_type_id" firestore:"vehicle_sub_type_id" query:"vehicle_sub_type_id" gorm:"-"`
	VehicleSubTypeName string `json:"vehicle_sub_type_name" firestore:"vehicle_sub_type_name" query:"vehicle_sub_type_name" gorm:"-"`
	VehicleModelID     string `json:"vehicle_model_id" firestore:"vehicle_model_id" query:"vehicle_model_id" gorm:"-"`
	VehicleModelName   string `json:"vehicle_model_name" firestore:"vehicle_model_name" query:"vehicle_model_name" gorm:"-"`
	IsOriginAirport    bool   `json:"is_origin_airport" firestore:"is_origin_airport" query:"is_origin_airport" gorm:"-"`
	DriverName         string `json:"driver_name" firestore:"driver_name" query:"driver_name" gorm:"-"`
	VehicleCode        string `json:"vehicle_code" firestore:"vehicle_code" query:"vehicle_code" gorm:"-"`
	DriverImage        string `json:"driver_image" firestore:"driver_image" query:"driver_image" gorm:"-"`
	VehicleImage       string `json:"vehicle_image" firestore:"vehicle_image" query:"vehicle_image" gorm:"-"`
	VehicleColorName   string `json:"vehicle_color_name" firestore:"vehicle_color_name" query:"vehicle_color_name" gorm:"-"`
	LicensePlate       string `json:"license_plate" firestore:"license_plate" query:"license_plate" gorm:"-"`
}
