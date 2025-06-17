package utils

import (
	"strings"

	"github.com/FourWD/aot.middleware/model"
	"github.com/FourWD/middleware/common"
)

func PublicNotiAppUser(group string, noti *model.NotiApp) error {
	topic := strings.ToUpper(common.App.Env) + "_USER_NOTI"
	return publicNotiApp(topic, group, noti)
}
