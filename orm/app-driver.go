package orm

type AppDriver struct {
	Slip
	SlipID             string `json:"slip_id" firestore:"slip_id" query:"slip_id" gorm:"-"`
	VehicleModelID     string `json:"vehicle_model_id" firestore:"vehicle_model_id" gorm:"-"`
	VehicleModelName   string `json:"vehicle_model_name" firestore:"vehicle_model_name" query:"vehicle_model_name" gorm:"-"`
	VehicleSubTypeID   string `json:"vehicle_sub_type_id" firestore:"vehicle_sub_type_id" gorm:"-"`
	VehicleSubTypeName string `json:"vehicle_sub_type_name" firestore:"vehicle_sub_type_name" gorm:"-"`
	IsOriginAirport    bool   `json:"is_origin_airport" firestore:"is_origin_airport" query:"is_origin_airport" gorm:"-"`
	DriverName         bool   `json:"driver_name" firestore:"driver_name" query:"driver_name" gorm:"-"`
	VehicleCode        bool   `json:"vehicle_code" firestore:"vehicle_code" query:"vehicle_code" gorm:"-"`
	DriverImage        bool   `json:"driver_image" firestore:"driver_image" query:"driver_image" gorm:"-"`
	VehicleImage       bool   `json:"vehicle_image" firestore:"vehicle_image" query:"vehicle_image" gorm:"-"`
	VehicleColorName   bool   `json:"vehicle_color_name" firestore:"vehicle_color_name" query:"vehicle_color_name" gorm:"-"`
	LicensePlate       bool   `json:"license_plate" firestore:"license_plate" query:"license_plate" gorm:"-"`
} //
