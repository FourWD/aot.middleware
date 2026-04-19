package utils

import (
	"strings"

	"github.com/FourWD/aot.middleware/model"
	"github.com/FourWD/middleware/infra"
)

func PublicNotiAppDriver(group string, noti *model.NotiApp) error {
	topic := strings.ToUpper(infra.AppInfo.Env) + "_DRIVER_NOTI"
	return publicNotiApp(topic, group, noti)
}
