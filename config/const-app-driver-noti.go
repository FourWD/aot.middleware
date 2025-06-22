package config

type appDriverNoti struct {
	NEW_JOB_APP      string
	NEW_JOB_DISPATCH string
}

var APP_DRIVER_NOTI = appDriverNoti{
	NEW_JOB_APP:      "NEW_JOB_APP",
	NEW_JOB_DISPATCH: "NEW_JOB_DISPATCH",
}
