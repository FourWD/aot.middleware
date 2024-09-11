package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/FourWD/aot.middleware/orm"
	"github.com/FourWD/middleware/common"
)

func GetAppUserFirebase(userID string) (orm.AppUser, error) {
	docPath := fmt.Sprintf("users/%s", userID)
	docRef := common.FirebaseClient.Doc(docPath)
	snap, err := docRef.Get(context.Background())
	if err != nil {
		return orm.AppUser{}, errors.New("cannot get app user")
	}

	var appUser orm.AppUser
	if err := snap.DataTo(&appUser); err != nil {
		return orm.AppUser{}, errors.New("cannot get app user")
	}

	return appUser, nil
}
