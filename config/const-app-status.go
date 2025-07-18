package config

type appUserStatus struct {
	AVAILABLE      string
	SEARCH_VEHICLE string
	WAIT_VEHICLE   string
	VEHICLE_ARRIVE string
	ON_THE_WAY     string
	DELIVER        string
}

type appDriverStatus struct {
	AVAILABLE    string
	ON_JOB       string
	JOB_MANUAL   string
	JOB_DISPATCH string
	JOB_CUSTOMER string
}

type appJobTypeID struct {
	JOB_MANUAL   string
	JOB_DISPATCH string
	JOB_SEARCH   string
}

type appSlipCancelTypeID struct {
	NOT_CANCEL       string
	CANCEL_BY_USER   string
	NOT_FOUND_DRIVER string
}

type appSlipTypeID struct {
	RETAIL     string
	CALLCENTER string
	FLEET      string
	WEBSITE    string
	APP        string
}

var APP_USER_STATUS = appUserStatus{
	AVAILABLE:      "00", // ว่าง
	SEARCH_VEHICLE: "01", // กำลังค้นหารถ
	WAIT_VEHICLE:   "02", // รถกำลังมารับคุณ
	VEHICLE_ARRIVE: "03", // รถมาถึงจุดรับคุณแล้ว
	ON_THE_WAY:     "04", // กำลังเดินทาง
	DELIVER:        "05", // ถึงปลายทาง
}

var APP_DRIVER_STATUS = appDriverStatus{
	AVAILABLE:    "00", // ว่าง
	ON_JOB:       "01", // มีงาน
	JOB_MANUAL:   "97", // คีย์งานตั๋วเอง
	JOB_DISPATCH: "98", // ส่วนกลางเรียกให้รับงาน (บังคับรับงาน)
	JOB_CUSTOMER: "99", // มีลูกค้าเรียกให้รับงาน (Broadcast)
}

var APP_JOB_DRIVER_TYPE = appJobTypeID{
	JOB_MANUAL:   "01", // คีย์งานตั๋วเอง
	JOB_DISPATCH: "02", // ส่วนกลางเรียกให้รับงาน (บังคับรับงาน)
	JOB_SEARCH:   "03", // ยูสเซอร์เรียก
}

var APP_SLIP_TYPE = appSlipTypeID{
	RETAIL:     "01",
	CALLCENTER: "02",
	FLEET:      "03",
	WEBSITE:    "04",
	APP:        "05",
}

var APP_SLIP_CANCEL_TYPE = appSlipCancelTypeID{
	NOT_CANCEL:       "00",
	CANCEL_BY_USER:   "01",
	NOT_FOUND_DRIVER: "03",
}
