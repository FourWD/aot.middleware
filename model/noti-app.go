package model

import "time"

type NotiApp struct {
	SlipID     string    `json:"slip_id" query:"slip_id"`
	VehicleID  string    `json:"vehicle_id" query:"vehicle_id"`
	DriverID   int       `json:"driver_id" query:"driver_id"`
	UserID     string    `json:"user_id" query:"user_id"`
	ActionTime time.Time `json:"action_time" query:"action_time"`
}
