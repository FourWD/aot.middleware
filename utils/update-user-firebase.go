package utils

import (
	"fmt"

	"github.com/FourWD/middleware/infra"
)

func UpdateUserFirebase(userID string, updateData map[string]interface{}) error {
	logFields := map[string]interface{}{
		"userID":     userID,
		"updateData": updateData,
	}
	infra.AppLog.Event("UpdateUserFirebase", logFields, userID)

	docPath := fmt.Sprintf("users/%s", userID)
	err := infra.FirebaseUpdate(infra.FirestoreClient, docPath, updateData)
	if err != nil {
		infra.AppLog.EventError(err, "UpdateUserFirebase Error", logFields, userID)
		return err
	}
	return nil
}
