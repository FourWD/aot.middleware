package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
)

func SendNotiToDriver(driverID, slipID, typeID string) {
	title := "NOTI-DRIVER"
	body := "TEST DRIVER"
	topic := fmt.Sprintf("DRIVER_%s", driverID)
	logData := map[string]interface{}{
		"topic":    topic,
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

	common.SendMessageToSubscriber(topic, title, body, data)
}
