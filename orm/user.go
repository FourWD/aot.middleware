package orm

import (
	orm "github.com/FourWD/middleware/model"
)

type User struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SourceID string `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`

	Username string `db:"username" json:"username" gorm:"type:varchar(50)"`

	Password    string `db:"password" json:"password" gorm:"type:varchar(100)"`
	Firstname   string `db:"firstname" json:"firstname" gorm:"type:varchar(50)"`
	Lastname    string `db:"lastname" json:"lastname" gorm:"type:varchar(50)"`
	Nickname    string `db:"nickname" json:"nickname" gorm:"type:varchar(50)"`
	Birthday    string `db:"birthday" json:"birthday" `
	Avatar      string `db:"avatar" json:"avatar" gorm:"type:varchar(100)"`
	Email       string `db:"email" json:"email" gorm:"type:varchar(50)"`
	Position    string `db:"position" json:"position" gorm:"type:varchar(100)"`
	UserGroupID string `db:"user_group_id" json:"user_group_id" gorm:"type:varchar(36)"`
}
