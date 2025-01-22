package orm

type Slip struct {
	ID string `json:"id" query:"id" firestore:"ID" gorm:"type:varchar(36);unique;"`
	SlipTemp
}
