package utils

import (
	"github.com/FourWD/aot.middleware/config"
	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func InitAppUser(userID string) error {
	appUser := new(orm.AppUser)
	appUser.AppUserStatusID = config.APP_USER_STATUS.AVAILABLE

	_, err := common.FirebaseClient.Collection("users").Doc(userID).Set(common.FirebaseCtx, appUser)
	return err
}