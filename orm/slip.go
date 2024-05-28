package orm

type Slip struct {
	SlipModel
	JobID string `json:"job_id" query:"job_id" gorm:"type:varchar(10);comment:'หมายเลขงาน'"` //หมายเลขงาน
}
