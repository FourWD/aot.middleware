package utils

import (
	"encoding/json"
	"errors"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/FourWD/aot.middleware/model"
	"github.com/FourWD/middleware/common"
)

func GetNotiApp(message *pubsub.Message) (model.NotiApp, error) {
	common.Log("GetNotiApp Start", map[string]interface{}{
		"message.id":   message.ID,
		"message.data": string(message.Data),
	}, message.ID)

	var noti model.NotiApp
	if err := json.Unmarshal(message.Data, &noti); err != nil {
		common.LogError(err.Error(), nil, message.ID)
		return noti, err
	}

	expire := noti.ActionTime.Add(5 * time.Minute)
	if time.Now().After(expire) {
		common.LogError("GetNotiApp expired", nil, message.ID)
		return noti, errors.New("noti-app expired")
	}

	return noti, nil
}
