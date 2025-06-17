package utils

import (
	"time"

	"github.com/FourWD/aot.middleware/model"
)

func IsNotiAppExpire(noti *model.NotiApp) bool {
	expire := noti.ActionTime.Add(5 * time.Minute)
	return time.Now().After(expire)
}
