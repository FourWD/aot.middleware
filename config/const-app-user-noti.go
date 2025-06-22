package config

type appUserNoti struct {
	PAYMENT_SUCCESS     string
	PAYMENT_FAIL        string
	SEARCH_VEHICLE      string
	WAIT_VEHICLE        string
	VEHICLE_ARRIVE      string
	ON_THE_WAY          string
	DELIVER             string
	CANCEL_BY_USER      string
	CANCEL_BY_DRIVER    string
	PROMOTION           string
	NOT_FOUND_DRIVER    string
	REVIEW_TRIP         string
	ALMOST_BOOKING_TIME string
}

var APP_USER_NOTI = appUserNoti{
	PAYMENT_SUCCESS:     "PAYMENT_SUCCESS",
	PAYMENT_FAIL:        "PAYMENT_FAIL",
	SEARCH_VEHICLE:      "SEARCH_VEHICLE",
	WAIT_VEHICLE:        "WAIT_VEHICLE",
	VEHICLE_ARRIVE:      "VEHICLE_ARRIVE",
	ON_THE_WAY:          "ON_THE_WAY",
	DELIVER:             "DELIVER",
	CANCEL_BY_USER:      "CANCEL_BY_USER",
	CANCEL_BY_DRIVER:    "CANCEL_BY_DRIVER",
	PROMOTION:           "PROMOTION",
	NOT_FOUND_DRIVER:    "NOT_FOUND_DRIVER",
	REVIEW_TRIP:         "REVIEW_TRIP",
	ALMOST_BOOKING_TIME: "ALMOST_BOOKING_TIME",
}
