package orm

type Slip struct {
	ID string `json:"id" query:"id" firestore:"ID" gorm:"type:varchar(36);unique;"`
	SlipTemp
	JobNo     string `json:"job_no" query:"job_no" firestore:"job_no" gorm:"type:varchar(36);unique"`
	RunningNo int    `json:"running_no" query:"running_no" firestore:"running_no" gorm:"primary_key;auto_increment;not_null"`
}
