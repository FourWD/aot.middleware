package utils

import (
	"github.com/FourWD/aot.middleware/config"
	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/infra"
)

func InitAppUser(userID string) error {
	if userID == "" {
		return nil
	}

	appUser := new(orm.AppUser)
	appUser.ID = userID

	appUser.AppUserStatusID = config.APP_USER_STATUS.AVAILABLE

	_, err := infra.FirestoreClient.Collection("users").Doc(userID).Set(infra.FirebaseCtx, appUser)
	return err
}
