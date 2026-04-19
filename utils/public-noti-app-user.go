package utils

import (
	"strings"

	"github.com/FourWD/aot.middleware/model"
	"github.com/FourWD/middleware/infra"
)

func PublicNotiAppUser(group string, noti *model.NotiApp) error {
	topic := strings.ToUpper(infra.AppInfo.Env) + "_USER_NOTI"
	return publicNotiApp(topic, group, noti)
}
