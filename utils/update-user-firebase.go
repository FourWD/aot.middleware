package utils

import (
	"fmt"

	"github.com/FourWD/middleware/common"
	"github.com/FourWD/middleware/infra"
)

func UpdateUserFirebase(userID string, updateData map[string]interface{}) error {
	logFields := map[string]interface{}{
		"userID":     userID,
		"updateData": updateData,
	}
	common.Log("UpdateUserFirebase", logFields, userID)

	docPath := fmt.Sprintf("users/%s", userID)
	err := common.FirebaseUpdate(infra.FirestoreClient, docPath, updateData)
	if err != nil {
		logFields["error"] = err
		common.LogError("UpdateUserFirebase Error", logFields, userID)
		return err
	}
	return nil
}
