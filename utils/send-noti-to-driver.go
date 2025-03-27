package utils

import (
	"github.com/FourWD/middleware/common"
)

func SendNotiToDriver(driverID, slipID, typeID string) {
	title := "NOTI-DRIVER"
	body := "TEST DRIVER"
	logData := map[string]interface{}{
		"driverID": driverID,
		"slipID":   slipID,
		"typeID":   typeID,
		"body":     body,
		"title":    title,
	}

	common.Log("SendNotiToDriver", logData, "")

	data := map[string]string{
		// "event_code": "NT001",
	}

	common.SendMessageToSubscriber(driverID, title, body, data)
}
