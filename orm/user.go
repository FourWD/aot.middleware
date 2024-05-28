package orm

import (
	"time"

	orm "github.com/FourWD/middleware/model"
)

type User struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SourceID    string    `json:"source_id" query:"source_id" gorm:"type:varchar(2);"` // ไหน table source
	Username    string    `json:"username" query:"username" gorm:"type:varchar(50)"`
	Password    string    `json:"password" query:"password" gorm:"type:varchar(100)"`
	Firstname   string    `json:"firstname" query:"firstname" gorm:"type:varchar(50)"`
	Lastname    string    `json:"lastname" query:"lastname" gorm:"type:varchar(50)"`
	Nickname    string    `json:"nickname" query:"nickname" gorm:"type:varchar(50)"`
	Birthday    time.Time `json:"birthday" query:"birthday" `
	Avatar      string    `json:"avatar" query:"avatar" gorm:"type:varchar(100)"`
	Email       string    `json:"email" query:"email" gorm:"type:varchar(100)"`
	Position    string    `json:"position" query:"position" gorm:"type:varchar(100)"`
	UserGroupID string    `json:"user_group_id" query:"user_group_id" gorm:"type:varchar(2)"`
}
